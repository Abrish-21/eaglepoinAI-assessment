package usecases

import (
    "time"
    "rate-limiter/entities"
)

// A simple in-memory storage (map). No interface needed.
var UserLimits = map[string]*entities.RateLimit{}

// Allow returns true if request is allowed.
func Allow(userID string) bool {
    limit, exists := UserLimits[userID]

    if !exists {
        // First time seeing user
        limit = &entities.RateLimit{
            UserID:       userID,
            Requests:     0,
            MaxRequests:  5,
            WindowStart:  time.Now(),
            WindowLength: 60 * time.Second,
        }
        UserLimits[userID] = limit
    }

    // Reset if 60s passed
    if time.Since(limit.WindowStart) > limit.WindowLength {
        limit.Requests = 0
        limit.WindowStart = time.Now()
    }

    // If limit hit → block
    if limit.Requests >= limit.MaxRequests {
        return false
    }

    // Count request + allow
    limit.Requests++
    return true
}
