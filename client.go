package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type client struct {
	conn     net.Conn
	username string
	room     *room
	commands chan<- command //all commands will go inside this channel and will be sent later to the server

}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			log.Println("unable to read the message")
			return
		}
		msg = strings.Trim(msg, "\r\n")
		//In order to build any specific command we have to extract it from message so we use split
		args := strings.Split(msg, " ")
		//now we use trimspace to exactly pick the right part from msg
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/username":
			//we will send the commands to the server and server will do the rest
			c.commands <- command{
				id:     cmd_username,
				client: c,
				args:   args}
		case "/join":
			c.commands <- command{
				id:     cmd_join,
				client: c,
				args:   args}
		case "/listrooms":
			c.commands <- command{
				id:     cmd_listRooms,
				client: c,
				args:   args}
		case "/msg":
			c.commands <- command{
				id:     cmd_msg,
				client: c,
				args:   args}
		case "/exit":
			c.commands <- command{
				id:     cmd_exit,
				client: c,
				args:   args}
		default:
			c.err(fmt.Errorf("unknown command %s", cmd))
		}
	}
}
func (c *client) err(err error) {
	c.conn.Write([]byte("err:" + err.Error() + "\n"))
}

func (c *client) msg(msg string) {
	if c.room != nil {
		c.room.lastActivity = time.Now()
	}
	c.conn.Write([]byte(">" + msg + "\n"))
}
