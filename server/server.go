/*package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
)

type Data struct {
	Body []byte
	Data []bye
}

func Listen() net.Listener {

	listener, e := net.Listen("tcp", "localhost:9000")
	if e != nil {
		fmt.Println(e.Error())
	}
	return listener

}
func main() {

	listener := Listen()
	publish := Data{Body: []byte("Publish"), Data: []byte("Hello")}
	buffer := new(bytes.Buffer)

	encoder := gob.NewEncoder(buffer)
	data := encoder.Encode(publish)

	conn, _ := listener.Accept()
	for {

		//fmt.Fprint(conn, "Hello")
		conn.Write(data)

		msg, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(msg)
	}

}
*/

package server

import (
	"encoding/gob"
	"log"
	"net"
)

type Data struct {
	Head []byte
	Body []byte
}

type Server struct {
	Addr      string
	Exchanges map[string]func(Data)
	Queues    map[string]chan Data
	Listener  net.Listener
}

func handleConnection(conn net.Conn) {

	for {
		dec := gob.NewDecoder(conn)

		d := &Data{}
		dec.Decode(d)

		if d != nil {
			log.Println(string(d.Body))
		}
	}
}
func (s *Server) Listen() {

	var e error
	var conn net.Conn

	s.Listener, e = net.Listen("tcp", s.Addr)

	if e != nil {

	}

	for {

		conn, e = s.Listener.Accept()

		go handleConnection(conn)
	}
}
