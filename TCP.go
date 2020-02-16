package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server(){
	ln, err := net.Listen("tcp", ":9090")
	if err != nil {
		return
	}
	for {
		conn, err := ln.Accept()
		if err !=nil {
			continue
		}
		var msg string
		err = gob.NewDecoder(conn).Decode(&msg)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Received:", msg)
		}
	}
}