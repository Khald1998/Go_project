package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", hello)

	err := http.ListenAndServe("localhost:9999", nil)

	if err != nil {
		log.Fatal(err)
	}
}
