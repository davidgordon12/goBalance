package main

type Request struct {
	host   string // 127.0.0.1
	method string // We will always assume HTTP/1.1
	agent  string //
	accept string // */*
}

func NewRequest(method string, host string) *Request {
	r := new(Request)
	r.host = host
	r.method = method
	r.agent = "curl"
	r.accept = "*"
	return r
}
