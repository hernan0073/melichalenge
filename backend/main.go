package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/time/rate"
)

var (
	// Default rate limits
	defaultRateLimit  = 1000 // Requests per minute for default IPs
	defaultBurstLimit = 5
	visitors          = make(map[string]*RateLimiter)
	client            = &http.Client{Timeout: 30 * time.Second}
)

type RateLimitConfig struct {
	Rate  int `json:"rate"`
	Burst int `json:"burst"`
}

type RateLimiter struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

func setRateLimit(w http.ResponseWriter, r *http.Request) {
	var config RateLimitConfig
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	defaultRateLimit = config.Rate
	defaultBurstLimit = config.Burst
	visitors = make(map[string]*RateLimiter)
	fmt.Fprintf(w, "Rate limit set to %d and burst limit to %d", defaultRateLimit, defaultBurstLimit)
}

func getRateLimiter(ip string, path string) *rate.Limiter {
	key := fmt.Sprintf("%s|%s", ip, path)
	if limiter, exists := visitors[key]; exists {
		limiter.lastSeen = time.Now()
		return limiter.limiter
	}

	limit := rate.Limit(defaultRateLimit)
	limiter := rate.NewLimiter(limit, defaultBurstLimit)
	visitors[key] = &RateLimiter{limiter, time.Now()}
	return limiter
}

func rateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := getRateLimiter(r.RemoteAddr, r.URL.Path)
		if !limiter.Allow() {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	targetURL := "https://api.mercadolibre.com" + r.RequestURI
	req, err := http.NewRequest(r.Method, targetURL, r.Body)
	if err != nil {
		log.Printf("Failed to create request for URL %s: %v", targetURL, err)
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	req.Header = r.Header
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to make request to URL %s: %v", targetURL, err)
		http.Error(w, "Failed to make request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		if key != "Transfer-Encoding" {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}
	}

	w.WriteHeader(resp.StatusCode)
	if _, err := io.Copy(w, resp.Body); err != nil {
		log.Printf("Failed to copy response body: %v", err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/rate-limit", setRateLimit).Methods("POST")
	r.Handle("/metrics", promhttp.Handler())
	r.PathPrefix("/").HandlerFunc(proxyHandler)
	r.Use(rateLimiterMiddleware)

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
