package main

import "github.com/autumn31/socket/server"

func main() {
	srv := server.New("8080")
	srv.Run()
}
