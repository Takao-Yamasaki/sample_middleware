package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	})

	http.Handle("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}