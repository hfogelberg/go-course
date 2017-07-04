package main

import (
	"io"
	"net/http"
)

func handleDog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Doggy")
}

func handleCat(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Kitty")
}

func main() {

	http.HandleFunc("/dog", handleDog)
	http.HandleFunc("/cat", handleCat)

	http.ListenAndServe(":8080", nil)
}
