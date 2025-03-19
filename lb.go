package main

import (
	"bufio"
	"fmt"
	"net"
)

func Serve() {
	listener, err := net.Listen("tcp", "localhost:80")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 80")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	fmt.Println("Received request from " + conn.LocalAddr().String())
	reader := bufio.NewReader(conn)
	var requestLines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil || line == "\r\n" {
			break
		}
		requestLines = append(requestLines, line)
	}
	for i := 0; i < len(requestLines); i += 1 {
		fmt.Print(requestLines[i])
	}
	defer conn.Close()
}
