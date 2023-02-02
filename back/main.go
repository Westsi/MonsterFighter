package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	initVars()
	monsterTester()
	initDB()

	http.HandleFunc("/", hello)
	http.HandleFunc("/getusers", getUsers)
	http.HandleFunc("/getuser", getUser)
	http.HandleFunc("/new/user", newUser)
	http.ListenAndServe("8000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
