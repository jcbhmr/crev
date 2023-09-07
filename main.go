package main

import (
	"log"
	"net/http"

	crev "jcbhmr.me/crev/pkg"
)

func main() {
	http.HandleFunc("/", crev.ServeHTTP)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
