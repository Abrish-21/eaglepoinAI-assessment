package main

import (
    "fmt"
    "log"
    "net/http"
    "rate-limiter/interfaces/middleware"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    userID := r.Header.Get("X-User-ID")
    fmt.Fprintf(w, "Hello, %s! Request allowed \n", userID)
}

func main() {
    mux := http.NewServeMux()
    mux.Handle("/", middleware.RateLimitMiddleware(http.HandlerFunc(helloHandler)))

    fmt.Println("Server running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
