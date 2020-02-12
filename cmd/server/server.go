package main

import (
	"fmt"
	"net"
)

var connSignal chan string = make(chan string)
var connections []net.Conn
var port string

func main() {

	port := ":7777"
	ln, _ := net.Listen("tcp", port)
	fmt.Printf("Server Status: Up\n")

	for {
		go session(ln)
		fmt.Println(<-connSignal)

	}

}

func session(ln net.Listener) {
	conn, _ := ln.Accept() //waits till it gets connection
	connSignal <- "Connection Established"

	connections = append(connections, conn)

	for {

		buf := make([]byte, 1024)
		conn.Read(buf)
		for _, c := range connections {
			c.Write(buf)
		}

		fmt.Print(string(buf))
	}
}
