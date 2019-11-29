package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
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
		logline(word, 20, 10)
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

func logline(fun string, line, sev uint) {
	vs := url.Values{}
	vs.Set("ServiceName", "webapp")
	vs.Set("FunctionName", fun)
	vs.Set("LineNumber", fmt.Sprintf("%d", line))
	vs.Set("Severity", fmt.Sprintf("%d", sev))
	go func(vs url.Values) {
		resp, err := http.PostForm("http://localhost:8090/log/formData", vs)
		fmt.Printf("resp:%v \nerr: %v", resp, err)
	}(vs)
}
