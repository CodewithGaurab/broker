package main

import (
	"broker/server"
)

func main() {

	s := server.Server{Addr: "localhost:9000"}

	s.Listen()

}
