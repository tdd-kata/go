package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Greet sends a personalised greeting to writer.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s\n", name)
}

// MyGreeterHandler says Hello, world over HTTP.
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "markruler")
}

func main() {
	// go run di.go
	// curl localhost:5000
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}
