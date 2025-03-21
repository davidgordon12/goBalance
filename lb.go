package main

import (
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	audit.info("/ request received from " + r.RemoteAddr)
	audit.info(r.Pattern + " " + r.Method + " " + r.Proto + " Host: " + r.Host + " User-Agent: " + r.UserAgent() + " Accept: ")
	_, err := http.Get("http://localhost:5000/")
	if err != nil {
		audit.error(err.Error())
		return
	}
}

func Serve() {
	go http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		audit.error(err.Error())
		return
	}
}
