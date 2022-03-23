package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to server
	con, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println("TCP Client: error on Dial:", err)
		return
	}
	for {
		// Get string from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("TCP Client: error on ReadString:", err)
			return
		}
		// send to server
		fmt.Fprintf(con, text+"\n")
		// wait response from server
		message, err := bufio.NewReader(con).ReadString('\n')
		if err != nil {
			fmt.Println("TCP Client: error on Server Response:", err)
			return
		}
		fmt.Print("Response from server: " + message)
	}
}
