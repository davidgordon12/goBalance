package main

import (
	"bufio"
	"net"
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
	buff := make([]byte, 1024)
	reader := bufio.NewReader(conn)
	size, err := reader.Read(buff)
	if err != nil {
		audit.error("Couldn't read request")
		return
	}
	audit.info((string)(buff[:size]))
	defer conn.Close()
}
