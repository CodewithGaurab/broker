package coordinator

import (
	"bytes"
	"encoding/gob"
	"io"
	"log"
	"net"
)

type Data struct {
	Data     []byte
	Exchange []byte
}

func Dial(addr string) (net.Conn, error) {

	connection, e := net.Dial("tcp", addr)

	if e != nil {
		log.Println(e.Error())
	}
	defer connection.Close()

	return &connection, e
}

func Publish(data Data, conn net.Conn) {

	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)
	encoded_data := encoder.Encode(data)
	conn.Write(encoded_data)

}

func Listen(addr string, conn net.Conn) {
	connection, e := net.Listen("tcp", addr)

	if e != nil {
		log.Println(e.Error())
	}
	defer connection.Close()

	go func() {
		for {
			conn, err := conn.
		}
	}()
}
