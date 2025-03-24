package main

type Request struct {
	host   string // 127.0.0.1
	method string // We will always assume HTTP/1.1
	path   string
	agent  string //
	accept string // */*
}

func NewRequest(method string, host string, path string) *Request {
	r := new(Request)
	r.host = host
	r.method = method
	r.path = path
	r.agent = "curl"
	r.accept = "*"
	return r
}
