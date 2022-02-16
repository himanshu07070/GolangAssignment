package main

import (
	"log"
	"net"
)

func main() {
	serv := newserver()
	go serv.run()
	go serv.checkActivity()
	listenAndServe, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("unable to start server : %s", err.Error())
	}
	defer listenAndServe.Close()

	log.Printf("Starting server at :8000")

	for {
		//for accepting infinite incoming connection
		conn, err := listenAndServe.Accept()
		if err != nil {
			log.Printf("unable to accept connection: %s ", err.Error())
			continue
		}
		go serv.newClient(conn)
	}
}
