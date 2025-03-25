package main

import (
	"net/http"
	"time"
)

/* Returns a list of healthy servers from the pool */
func StartHealthcheckService() {
	for {
		time.Sleep(60 * time.Second)
		audit.info("Performing healthcheck")
		performHealthcheck()
	}
}

func performHealthcheck() {
	for i := 0; i < len(serverPool); i += 1 {
		resp, err := http.Get(serverPool[i])
		if err != nil || resp.Status != "200 OK" {
			audit.info("Removing unhealthy server from pool " + serverPool[i])
			addServer(serverPool[i], &unhealthyServers)
			popServer(serverPool[i], &serverPool)
			i -= 1
			if len(serverPool) == 0 {
				break
			}
		}
	}

	checkUnhealthyServers()
}

func checkUnhealthyServers() {
	for i := 0; i < len(unhealthyServers); i += 1 {
		resp, err := http.Get(unhealthyServers[i])
		if err == nil && resp.Status == "200 OK" {
			audit.info("Adding healthy server back to pool " + unhealthyServers[i])
			addServer(unhealthyServers[i], &serverPool)
			popServer(unhealthyServers[i], &unhealthyServers)
		}
	}
}
