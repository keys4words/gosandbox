package main

import (
	"bufio"
	"fmt"
	"net"
)

func handle(con net.Conn) {

}

func main() {
	fmt.Println("Server listening on 8081...")
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("TCP Server: error on listening:", err)
		return
	}
	defer ln.Close()

	con, _ := ln.Accept()
	for {
		//Listen all messages with line sep \n
		message, _ := bufio.NewReader(con).ReadString('\n')
		// Print message from client
		fmt.Print("Message from Client: ", string(message))

		con.Write([]byte("Received: OK\n"))
	}
}
