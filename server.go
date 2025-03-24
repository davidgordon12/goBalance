package main

import (
	"fmt"
	"io"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got request")
	io.WriteString(w, "Hello from server")
}

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", getRoot)
	err := http.ListenAndServe(":5000", mux)
	if err != nil {
		fmt.Println(err)
		return
	}
}
