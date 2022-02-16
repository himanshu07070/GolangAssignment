package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type server struct {
	rooms    map[string]*room
	commands chan command // it is a single channel where all commands from all clients will be sent to
}

func newserver() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Printf("new client has connected: %s ", conn.RemoteAddr().String())
	//now let's initialize client
	c := &client{
		conn:     conn,
		username: "Anonymous",
		commands: s.commands,
	}
	c.readInput()
}

func (s *server) run() {
	for cmd := range s.commands {
		//now here we are doing the logic part and switching on id
		switch cmd.id {
		case cmd_username:
			s.username(cmd.client, cmd.args)
		case cmd_join:
			s.join(cmd.client, cmd.args)
		case cmd_listRooms:
			s.listRooms(cmd.client)
		case cmd_msg:
			s.msg(cmd.client, cmd.args)
		case cmd_exit:
			s.exit(cmd.client, cmd.args)

		}
	}

}
func (s *server) username(c *client, args []string) {
	c.username = args[1]
	c.msg(fmt.Sprintf("Hi, %s", c.username))
}

func (s *server) join(c *client, args []string) {
	roomName := args[1]
	//first we check if any room is available or not
	r, ok := s.rooms[roomName]
	if !ok {
		//if not then create one
		r = &room{
			name:         roomName,
			members:      make(map[net.Addr]*client),
			lastActivity: time.Now(),
		}
		s.rooms[roomName] = r
	}
	//if the room is present then add the member directly
	if len(s.rooms[roomName].members) > 5 {
		log.Print("room is already full")
		return
	}
	r.members[c.conn.RemoteAddr()] = c
	//now we are marking on the client side the current updated room
	c.room = r
	r.broadcast(c, fmt.Sprintf("%s has joined the room", c.username))
	c.msg(fmt.Sprintf("%s, welcome to the group", c.username))
}
func (s *server) msg(c *client, args []string) {
	if c.room == nil {
		c.err(errors.New("you must join the room first"))
		return
	}
	c.room.broadcast(c, c.username+": "+strings.Join(args[1:], " ")+"\n")
}
func (s *server) exit(c *client, args []string) {

	log.Printf("client has disconnected: %s", c.conn.RemoteAddr().String())
	s.quitCurrentRoom(c)
	c.msg("bye")
	c.conn.Close()
}

func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		delete(c.room.members, c.conn.RemoteAddr())
		c.room.broadcast(c, fmt.Sprintf("%s has left the room", c.username))
	}

}

func (s *server) listRooms(c *client) {
	if len(s.rooms) == 0 {
		c.msg("no existing rooms, please create one using /join <roomname>")
		return
	}
	var allRooms string
	for roomname := range s.rooms {
		allRooms += roomname + "\n"
	}
	c.msg(allRooms)
}

func (s *server) checkActivity() {
	for {
		for _, room := range s.rooms {
			t := time.Now()
			if room.lastActivity.Add(30 * time.Second).Before(t) {
				for _, member := range room.members {
					member.msg("no activity in past 30 seconds , so closing the group in 10 seconds")
				}
				time.Sleep(10 * time.Second)
				for address, member := range room.members {
					delete(room.members, address)
					member.msg("closing the group : " + room.name)

				}
				delete(s.rooms, room.name)
			}
		}
	}
}
