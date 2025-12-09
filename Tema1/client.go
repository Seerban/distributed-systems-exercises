package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)

	fmt.Print("Nume: ")
	text, _ := reader.ReadString('\n')
	conn.Write([]byte(text))

	// conexiune server
	reply, _ := serverReader.ReadString('\n')
	fmt.Println("Server:", reply)

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		conn.Write([]byte(text))

		// "primit request"
		reply, _ := serverReader.ReadString('\n')
		fmt.Print("Server: ", reply)

		// raspuns
		reply2, _ := serverReader.ReadString('\n')
		fmt.Print("Server: ", reply2)
	}
}
