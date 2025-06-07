// client
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// connection to server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("[!] Error connection: ", err)
	}

	defer conn.Close()

	fmt.Println("[+] Connection estimatet. Type \"exit\" to quit")

	reader := bufio.NewReader(os.Stdin)

	for {
		// reading user's input
		fmt.Print("> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("[!] Error reading input: ", err)
			return
		}

		// sending message to server
		_, err = conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Error sending message: ", err)
			return
		}

		// if input is "exit" - close connection
		if strings.TrimSpace(text) == "exit" {
			fmt.Println("[*] Closing connection")
			return
		}

		// reading response from server
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("[!] Error reading response: ", err)
			return
		}

		fmt.Println("Server response: ", response)

	}

}
