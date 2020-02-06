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
	for {
		conn, err = net.Dial("tcp", "localhost:6060")
		if err == nil { //
			break
		}
	}
	fmt.Println("Client Status: Looking for Server at port 6060.")
	go Listenter(conn)

	Writer(conn)
}

func Listenter(conn net.Conn) {

	for {

		buf := make([]byte, 1024)
		conn.Read(buf)

		fmt.Print(string(buf))
	}

}

func Writer(conn net.Conn) {
	conn.Write([]byte("\nConnection Established\n"))

	for {

		r := bufio.NewReader(os.Stdin)
		text, _ := r.ReadString('\n')
		// fmt.Print(conn.LocalAddr().String)
		// fmt.Print(">>>>  ")
		conn.Write([]byte(conn.LocalAddr().String() + " >>>> " + text))
	}

}
