package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Server struct {
	port string
}

func New(port string) *Server {
	return &Server{port}
}

func (srv *Server) Run() {
	ln, err := net.Listen("tcp", ":"+srv.port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Println("Waiting Connections")
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Connecting with ", conn.LocalAddr(), conn.RemoteAddr())
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for {
		success := scanner.Scan()
		if !success {
			fmt.Println(scanner.Err())
		} else {
			fmt.Println(string(scanner.Bytes()))
		}
	}
}
