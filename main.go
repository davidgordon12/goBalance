package main

var audit *Audit

func main() {
	// Init audit
	audit = NewAudit()
	audit.addFile("logs.txt")

	// Init server pool
	addServer("http://127.0.0.1:5000", &serverPool)
	addServer("http://127.0.0.1:5001", &serverPool)
	addServer("http://127.0.0.1:5002", &serverPool)
	addServer("http://127.0.0.1:5003", &serverPool)

	go StartHealthcheckService()

	Serve()
}
