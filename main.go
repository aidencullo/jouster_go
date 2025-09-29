package main

import (
	"log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	
	log.Printf("Request: %s %s %s", r.Method, r.URL.Path, r.Proto)
	log.Printf("Remote address: %s", r.RemoteAddr)
	log.Printf("User-Agent: %s", r.UserAgent())

	message := "Hello from the server!"
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	
	if _, err := w.Write([]byte(message + "\n")); err != nil {
		log.Printf("Error writing response: %v", err)
	}

	duration := time.Since(start)
	log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, duration)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.Println("Starting server on :8080...")

	http.HandleFunc("/", hello)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
