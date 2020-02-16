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

func client() {
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		return
	}
	msg := "Helllo Go"
	err = gob.NewEncoder(conn).Encode(msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sending:", msg)
	}
	conn.Close
}

func main(){
	go server()
	go client()

	var input string
	fmt.Scan(&input)
}