package main

import (
	"bufio"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	audit.info("/ request received from " + r.RemoteAddr)
	audit.info(r.Pattern + " " + r.Method + " " + r.Proto + " Host: " + r.Host + " User-Agent: " + r.UserAgent())
	result, err := http.Get("http://localhost:5000/")
	if err != nil {
		audit.error(err.Error())
		return
	}
	defer result.Body.Close()
	buffer := make([]byte, 1024)
	reader := bufio.NewReader(result.Body)
	size, err := reader.Read(buffer)
	if err != nil {
		audit.error(err.Error())
		return
	}
	w.Write(buffer[:size])
}

func Serve() {
	go http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		audit.error(err.Error())
		return
	}

	audit.info("Server listening on port 3000")
}
