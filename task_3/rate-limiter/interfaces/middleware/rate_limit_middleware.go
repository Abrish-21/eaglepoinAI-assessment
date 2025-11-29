package middleware

import (
    "net/http"
    "rate-limiter/usecases"
)

// RateLimitMiddleware checks each request against the rate limiter before allowing it to proceed.
func RateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Identify the user from the header
        userID := r.Header.Get("X-User-ID")
        if userID == "" {
            // If no user ID is provided, reject the request
            http.Error(w, "Missing X-User-ID header", http.StatusBadRequest)
            return
        }

        // Check if this user is allowed to make a request
        if !usecases.Allow(userID) {
            http.Error(w, "Rate limit exceeded. Try again later.", http.StatusTooManyRequests)
            return
        }

        // User is allowed, pass request to the actual handler
        next.ServeHTTP(w, r)
    })
}
