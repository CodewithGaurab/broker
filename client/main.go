package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Data struct {
	Head []byte
	Body []byte
}

func Dial() net.Conn {

	listener, e := net.Dial("tcp", "localhost:9000")
	if e != nil {
		fmt.Println(e.Error())
	}
	return listener

}
func main() {

	conn := Dial()
	defer conn.Close()
	d := &Data{Body: []byte("Sup dude!")}
	for {
		enc := gob.NewEncoder(conn)
		enc.Encode(d)
	}

}
