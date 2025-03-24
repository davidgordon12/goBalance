package main

import (
	"bufio"
	"net"
	"net/http"
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
	defer conn.Close()

	audit.info("Received request from " + conn.LocalAddr().String())

	reader := bufio.NewReader(conn)
	buff := make([]byte, 1024)
	size, err := reader.Read(buff)
	if err != nil {
		audit.error("Couldn't read request " + err.Error())
		return
	}

	audit.info((string)(buff[:size]))

	resp := forwardRequest()

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	conn.Write(resp)
}

func forwardRequest() []byte {
	resp, err := http.Get("http://localhost:5000")
	if err != nil {
		audit.error("Couldn't forward request " + err.Error())
		return nil
	}

	buff := make([]byte, 32000)
	httpReader := bufio.NewReader(resp.Body)
	size, err := httpReader.Read(buff)
	if err != nil {
		audit.error("Couldn't read request " + err.Error())
		return nil
	}

	return buff[:size]
}
