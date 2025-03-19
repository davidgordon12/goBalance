package main

import (
	"bufio"
	"net"
	"strings"
)

func Serve() {
	listener, err := net.Listen("tcp", "localhost:80")
	if err != nil {
		audit.error(err.Error())
		return
	}
	defer listener.Close()

	audit.info("Server listening on port 80")

	for {
		conn, err := listener.Accept()
		if err != nil {
			audit.error(err.Error())
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	audit.info("Received request from " + conn.LocalAddr().String())
	reader := bufio.NewReader(conn)
	var requestLines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil || line == "\r\n" {
			break
		}
		requestLines = append(requestLines, line)
	}
	audit.info(strings.Join(requestLines, "\n"))
	defer conn.Close()
}
