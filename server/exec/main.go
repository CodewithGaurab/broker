package main

import (
	"broker/server"
	"broker/utils"
	"log"
)

func direct(queues map[string]chan utils.Data, d utils.Data) {

	// log.Println(d)
	for _, val := range queues {
		// log.Println(val)
		val <- d
	}
	return
}
func main() {
	q := make(chan utils.Data)
	e := server.NewExchange(direct)
	e.BindQueue("q1", q)

	s := server.NewServer("localhost:9000")
	s.Exchanges["direct"] = e
	go s.Listen()
	//log.Println("Im here")
	for {
		log.Println("ok", <-q)
	}
}
