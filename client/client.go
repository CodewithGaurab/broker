package main

import (
	"broker/utils"
	"encoding/gob"
	"log"
	"net"
)

func Dial() net.Conn {

	listener, e := net.Dial("tcp", "localhost:9000")
	if e != nil {
		log.Println(e.Error())
	}
	return listener

}
func main() {

	conn := Dial()
	defer conn.Close()

	//reqd := &utils.Data{Head: []byte("q1"), Body: []byte("Give me data"), Exchange: "direct", Type: 0}
	for {
		//d3 := &utils.Data{}
		d := &utils.Data{Body: []byte("Sup dude!"), Exchange: "direct", Type: 1}
		enc := gob.NewEncoder(conn)
		//enc2 := gob.NewEncoder(conn)
		enc.Encode(d)

		//enc2.Encode(reqd)

		//enc3 := gob.NewDecoder(conn)
		//enc3.Decode(&d3)

	}

}
