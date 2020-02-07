package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	var conn net.Conn
	var err error
	port := ":6060"
	for {

		conn, err = net.Dial("tcp", port)
		if err == nil {
			break
		}
	}
	fmt.Println("Client Status: Looking for Server at port " + port)
	go listener(conn)

	writer(conn)
}

func listener(conn net.Conn) {

	for {

		buf := make([]byte, 1024)
		conn.Read(buf)

		fmt.Print(string(buf))
	}

}

func writer(conn net.Conn) {
	conn.Write([]byte("\nNew User!\n" + conn.LocalAddr().String() + "\n"))

	for {

		r := bufio.NewReader(os.Stdin)
		text, _ := r.ReadString('\n')

		conn.Write([]byte("  " + conn.LocalAddr().String() + " >>>> " + text))
	}

}
