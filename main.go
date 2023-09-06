package main

import (
	"os"
	"log"
)

func main() {
	var registryHost string
	if len(os.Args) >= 2 {
		registryHost = os.Args[1]
	}
	if registryHost == "" {
		registryHost = "ghcr.io"
	}

	var repoPrefix string
	if len(os.Args) >= 3 {
		repoPrefix = os.Args[2]
	}

	log.Fatal(ListenAndServe(":8080", registryHost, repoPrefix))
}
