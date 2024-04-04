package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handler)

	log.Print("Starting HTTP server on port 80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}