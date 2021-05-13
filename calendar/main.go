package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello v1!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
