package main

var unhealthyServers []string
var serverPool []string

func addServer(address string, pool *[]string) {
	// Sanitize first when done testing
	*pool = append(*pool, address)
}

func popServer(address string, pool *[]string) {
	// Sanitize first when done testing
	idx := -1
	for i := 0; i < len(*pool); i += 1 {
		if (*pool)[i] == address {
			idx = i
		}
	}
	if idx > 0 {
		*pool = append((*pool)[:idx], (*pool)[idx+1:]...)
	}
}
