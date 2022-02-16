package main

import (
	"net"
	"time"
)

type room struct {
	name         string
	members      map[net.Addr]*client //using remote address of clinet as map key and then value is pointer to the client
	lastActivity time.Time
}

func (r *room) broadcast(sender *client, msg string) {

	for addr, m := range r.members {
		if addr != sender.conn.RemoteAddr() {
			//send message to everyone except sender itself
			m.msg(msg)
		}
	}
}
