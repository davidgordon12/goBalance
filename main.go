package main

var audit *Audit

func main() {
	// Setup audit library before starting
	audit = NewAudit()
	Serve()
}
