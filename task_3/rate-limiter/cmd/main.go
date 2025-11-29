package main

import (
    "fmt"
    "log"
    "net/http"
    "rate-limiter/interfaces/middleware"
)

// Simple handler: just greets the user. Rate limiting is handled by middleware.
func helloHandler(w http.ResponseWriter, r *http.Request) {
    userID := r.Header.Get("X-User-ID")
    fmt.Fprintf(w, "Hello, %s! Request allowed \n", userID)
}

func main() {
    mux := http.NewServeMux()

    // Wrap our handler with rate limiter so each request is checked
    mux.Handle("/", middleware.RateLimitMiddleware(http.HandlerFunc(helloHandler)))

    fmt.Println("Server running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
