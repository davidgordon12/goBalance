package main

var audit *Audit

func main() {
	// Init audit
	audit = NewAudit()
	audit.addFile("logs.txt")
	Serve()
}
