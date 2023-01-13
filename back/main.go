package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	
	// http.HandleFunc("/", hello)
	// http.ListenAndServe("8000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
