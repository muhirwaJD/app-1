package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getHealth(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK")
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
    data, _ := os.ReadFile("./version.txt")
    hostname, _ := os.Hostname()

    w.Header().Set("Content-Type", "text/plain")
    fmt.Fprintf(w, "Version: %s\nHostname: %s\nTime: %s\n",
        string(data), hostname, time.Now().Format(time.RFC3339))
}

func main() {
	http.HandleFunc("/healthz", getHealth)
	http.HandleFunc("/version", fileHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}