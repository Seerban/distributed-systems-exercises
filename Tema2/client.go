package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)

	// login
	fmt.Print("Nume: ")
	name, _ := reader.ReadString('\n')
	conn.Write([]byte(name))

	serverReply, _ := serverReader.ReadString('\n')
	fmt.Println("Server:", strings.TrimSpace(serverReply))

	for {
		for {
			fmt.Print("> ")
			line, _ := reader.ReadString('\n')
			conn.Write([]byte(line))

			// empty line ends request
			if strings.TrimSpace(line) == "" {
				break
			}
		}
		// wait for server responses
		reply1, _ := serverReader.ReadString('\n')
		fmt.Println("Server:", strings.TrimSpace(reply1))
	}
}
