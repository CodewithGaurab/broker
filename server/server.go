package server

import (
	"broker/utils"
	"encoding/gob"
	"log"
	"net"
)

type Exchange struct {
	Queues map[string]chan utils.Data
	Call   func(map[string]chan utils.Data, utils.Data)
}

func NewExchange(c func(map[string]chan utils.Data, utils.Data)) *Exchange {
	s := Exchange{}
	s.Call = c
	s.Queues = make(map[string]chan utils.Data)

	return &s
}

func (e *Exchange) BindQueue(name string, channel chan utils.Data) {
	e.Queues[name] = channel
}

type Server struct {
	Addr      string
	Exchanges map[string]*Exchange
	Listener  net.Listener
}

func NewServer(addr string) *Server {

	s := Server{}
	s.Exchanges = make(map[string]*Exchange)
	s.Addr = addr

	return &s
}

func (s *Server) HandleConnection(conn net.Conn) {

	for {
		dec := gob.NewDecoder(conn)

		d := &utils.Data{}
		dec.Decode(d)
		/*
			if d != nil {
				log.Println(string(d.Body))
			}
		*/

		if d.Type == 1 && s.Exchanges[d.Exchange] != nil && d != nil {
			//enc := gob.NewEncoder(conn)
			//s.Exchanges[d.Exchange].Queues[string(d.Head)] <- *d
			//enc.Encode(d)
			go s.Exchanges[d.Exchange].Call(s.Exchanges[d.Exchange].Queues, *d)

		}

		if d.Type == 0 {

			for _, ex := range s.Exchanges {
				if ex.Queues[string(d.Head)] != nil {
					data := <-ex.Queues[string(d.Head)]
					enc := gob.NewEncoder(conn)

					enc.Encode(data)
				}
			}
		}

	}
}
func (s *Server) Listen() {

	var e error
	var conn net.Conn

	s.Listener, e = net.Listen("tcp", s.Addr)

	if e != nil {
		log.Println(e.Error())
	}

	for {

		conn, e = s.Listener.Accept()

		if conn != nil {
			go s.HandleConnection(conn)
		}
	}
}
