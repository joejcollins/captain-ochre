package main

import (
	"log"
	"net/http"
)

/*
Define a home handler function which writes a byte slice containing
some arbitrary string in the response body
*/
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Here"))
}

func showCSV(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Show CSV"))
}

func createCSV(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
		w.Write([]byte("Method Not Allowed"))
		w.WriteHeader(405)
        return
    }
    w.Write([]byte("Create CSV"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/status", showCSV)
	mux.HandleFunc("/status/update", createCSV)

	log.Println("Starting server on :http-alt")
	err := http.ListenAndServe(":http-alt", mux)
	log.Fatal(err)
}
