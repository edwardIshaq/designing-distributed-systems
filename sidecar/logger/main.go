package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

const maxMultipartFormSize = 10 * 1024 * 1024
const loggingPort = ":8090"

var hostName string

func init() {
	if name, err := os.Hostname(); err != nil {
		log.Printf("Error getting hostname: %s", err.Error())
	} else {
		hostName = name
		log.Printf("Running on Host: %s", hostName)
	}
}

func main() {
	debugInfo := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "This is the logger speaking\n")
	}

	http.HandleFunc("/", debugInfo)
	http.HandleFunc("/log/raw", logRaw)
	http.HandleFunc("/log/urlEncoded", logFormURLEncoded)
	http.HandleFunc("/log/formData", logFormURLEncoded)

	log.Fatal(http.ListenAndServe(loggingPort, nil))
}

//x-www-form-urlencoded POST see URLEncoded.png
func logFormURLEncoded(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(maxMultipartFormSize)
	var entry LogEntry
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Please send POST requests only")
		return
	}

	bs := []byte{}
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Couldn't parse POST")
		return
	}
	fmt.Printf("Body: %s\nForm: %v", string(bs), r.Form)

	entry.ServiceName = r.Form.Get("ServiceName")
	entry.FunctionName = r.Form.Get("FunctionName")
	num, err := strconv.ParseUint(r.Form.Get("LineNumber"), 10, 32)
	if err == nil {
		entry.LineNumber = uint(num)
	}
	num, err = strconv.ParseUint(r.Form.Get("Severity"), 10, 32)
	if err == nil {
		entry.Severity = uint(num)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(entry)
}

//raw json POST
func logRaw(w http.ResponseWriter, r *http.Request) {
	var entry LogEntry

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Please send POST requests only")
		return
	}

	bs := []byte{}
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Couldn't parse POST")
		return
	}
	fmt.Printf("Body: %s\nForm: %v", string(bs), r.Form)

	buff := bytes.NewBuffer(bs)
	if err := json.NewDecoder(buff).Decode(&entry); err != nil && err != io.EOF {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Couldn't parse POST")
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(entry)
}

// LogEntry is the basic type of log
type LogEntry struct {
	ServiceName  string `json:"ServiceName,omitempty"`
	Module       string `json:"Module,omitempty"`
	FunctionName string `json:"FunctionName,omitempty"`
	LineNumber   uint   `json:"LineNumber,omitempty"`
	Severity     uint   `json:"Severity,omitempty"`
}

func reverse(word string) string {
	flipped := ""
	for _, c := range word {
		flipped = fmt.Sprintf("%c%s", c, flipped)
	}
	return flipped
}
