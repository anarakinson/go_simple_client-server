// server
package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {

	// starting TCP-server on 8080 port
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("[!] Error listening: ", err)
		return
	}

	defer listener.Close()

	fmt.Println("[*] Server is listening on port :8080")

	for {
		// accepting connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("[!] Error connection: ", err)
			continue
		}

		// handling connection in separate goroutine
		go HandleConnection(conn)
	}

}

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		// read message from client
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("[!] Error reading: ", err)
			return
		}

		// strip spaces
		message = strings.TrimSpace(message)
		fmt.Printf("[+] Received: %s \n", message)

		// if received "exit" - close connection
		if message == "exit" {
			fmt.Println("[*] Close connection")
		}

		// sending responce
		response := "Your message is very important for us!\n"
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("[!] Error writing: ", err)
			return
		}
	}

}
