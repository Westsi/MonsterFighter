package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	initVars()
	// monsterTester()
	fmt.Println(genNewID(false, "I am test monster", "8ed3bcab635e02b5489a029d45aaefaa0bdda63b23b09c674a493e91e7453a15"))
	getParentObjectFromID("test.json")
	initDB()

	// http.HandleFunc("/", hello)
	// http.ListenAndServe("8000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
