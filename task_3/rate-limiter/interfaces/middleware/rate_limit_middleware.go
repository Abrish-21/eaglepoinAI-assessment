package middleware

import (
    "net/http"
    "rate-limiter/usecases"
)

func RateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        userID := r.Header.Get("X-User-ID")
        if userID == "" {
            http.Error(w, "Missing X-User-ID header", http.StatusBadRequest)
            return
        }

        if !usecases.Allow(userID) {
            http.Error(w, "Rate limit exceeded. Try again later.", http.StatusTooManyRequests)
            return
        }

        next.ServeHTTP(w, r)
    })
}
