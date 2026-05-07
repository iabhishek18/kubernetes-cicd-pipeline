package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Version   string `json:"version"`
	Hostname  string `json:"hostname"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "1.0.0"
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		resp := HealthResponse{Status: "healthy", Timestamp: time.Now().UTC().Format(time.RFC3339), Version: version, Hostname: hostname}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ready")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Kubernetes CI/CD Demo App v%s", version)
	})

	log.Printf("Server starting on port %s (version %s)", port, version)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
