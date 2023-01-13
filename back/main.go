package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	initVars()
	bm := createMonster([]Monster{WormMonster, TrollMonster})
	fmt.Println(bm.Name)
	fmt.Println(bm.Rarity)
	fmt.Println(bm.Health)
	fmt.Println(bm.Generation)

	// http.HandleFunc("/", hello)
	// http.ListenAndServe("8000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
