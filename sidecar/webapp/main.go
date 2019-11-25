package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		word := r.Form.Get("word")
		response := fmt.Sprintf("Reversing %s => %s\n", word, reverse(word))
		io.WriteString(w, response)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/api", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func reverse(word string) string {
	flipped := ""
	for _, c := range word {
		flipped = fmt.Sprintf("%c%s", c, flipped)
	}
	return flipped
}
