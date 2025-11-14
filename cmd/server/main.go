package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BlackDadd77/3V30OStudios/internal/config"
)

func main() {
	cfg := config.Load()

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/health", handleHealth)

	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	fmt.Printf("Starting 3V30OStudios server on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message":"Welcome to 3V30OStudios API","status":"running"}`)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"healthy"}`)
}
