package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"sync/atomic"
	"time"
)

var (
	requestCount uint64
	startTime    = time.Now()
	version      = getEnv("APP_VERSION", "2.0.0")
	ready        = true
)

type HealthResponse struct {
	Status       string  `json:"status"`
	Version      string  `json:"version"`
	Uptime       string  `json:"uptime"`
	Hostname     string  `json:"hostname"`
	GoVersion    string  `json:"go_version"`
	NumGoroutine int     `json:"goroutines"`
	MemoryMB     float64 `json:"memory_mb"`
	Requests     uint64  `json:"total_requests"`
}

type MetricsResponse struct {
	Uptime       float64 `json:"uptime_seconds"`
	Requests     uint64  `json:"total_requests"`
	Goroutines   int     `json:"goroutines"`
	HeapMB       float64 `json:"heap_mb"`
	StackMB      float64 `json:"stack_mb"`
	NumGC        uint32  `json:"gc_cycles"`
}

func main() {
	port := getEnv("PORT", "8080")

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/ready", readyHandler)
	http.HandleFunc("/metrics", metricsHandler)
	http.HandleFunc("/", rootHandler)

	log.Printf("[Server] Starting on :%s (version %s)", port, version)
	log.Fatal(http.ListenAndServe(":"+port, logMiddleware(http.DefaultServeMux)))
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint64(&requestCount, 1)
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %s %s (%s)", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	})
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	hostname, _ := os.Hostname()

	resp := HealthResponse{
		Status:       "healthy",
		Version:      version,
		Uptime:       time.Since(startTime).Round(time.Second).String(),
		Hostname:     hostname,
		GoVersion:    runtime.Version(),
		NumGoroutine: runtime.NumGoroutine(),
		MemoryMB:     float64(m.Alloc) / 1024 / 1024,
		Requests:     atomic.LoadUint64(&requestCount),
	}
	writeJSON(w, http.StatusOK, resp)
}

func readyHandler(w http.ResponseWriter, _ *http.Request) {
	if ready {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ready")
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprint(w, "not ready")
	}
}

func metricsHandler(w http.ResponseWriter, _ *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	resp := MetricsResponse{
		Uptime:     time.Since(startTime).Seconds(),
		Requests:   atomic.LoadUint64(&requestCount),
		Goroutines: runtime.NumGoroutine(),
		HeapMB:     float64(m.HeapAlloc) / 1024 / 1024,
		StackMB:    float64(m.StackInuse) / 1024 / 1024,
		NumGC:      m.NumGC,
	}
	writeJSON(w, http.StatusOK, resp)
}

func rootHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"service": "k8s-demo", "version": version, "status": "running"})
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
