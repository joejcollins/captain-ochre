package main

import (
	"log"
	"net/http"
)

// home handler requests to the root.
func home(w http.ResponseWriter, r *http.Request) {
	// Don't show paths that we don't have
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	path := r.URL.Path
	log.Println(path)
	w.Write([]byte("Hello from Here"))
}

// showCSV should show a selected CSV file from the processed directory
func showCSV(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Show CSV"))
}

// createCSV just trying out a POST
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
