package main

import (
	"net"
	"log"
	"broker/utils"
	"encoding/gob"
)


func main(){

	conn,e := net.Dial("tcp","localhost:9000")
	d := &utils.Data{Body : []byte("Hola!"),Type : 1, Exchange : "direct"}
	if e != nil{
		log.Println(e.Error())	
	}
	for{
		enc := gob.NewEncoder(conn)
		enc.Encode(d)
	}

}
