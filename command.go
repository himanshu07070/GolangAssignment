package main

type commandId int

const (
	cmd_username commandId = iota
	cmd_join
	cmd_listRooms
	cmd_msg
	cmd_exit
)

type command struct {
	id     commandId
	client *client
	args   []string
}
