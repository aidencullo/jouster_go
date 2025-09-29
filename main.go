package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	message := "Hello from the server!"
	fmt.Println(message)
	fmt.Fprintln(w, "Hello, world")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
