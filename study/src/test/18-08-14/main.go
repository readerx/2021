package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// connect to this socket
	conn, _ := net.Dial("tcp6", "[2001:19f0:9002:1780:5400:1ff:fea1:ec51]:8080")

	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text+"\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
