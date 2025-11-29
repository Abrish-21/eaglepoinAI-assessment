package usecases

import (
    "time"
    "rate-limiter/entities"
)

// In-memory storage to track requests per user
var UserLimits = map[string]*entities.RateLimit{}

// Allow checks if a user can make a request based on rate limits
func Allow(userID string) bool {
    limit, ok := UserLimits[userID]
    if !ok {
        // First time seeing this user → create a rate limit record
        limit = &entities.RateLimit{
            Requests:     0,
            MaxRequests:  5,                  // max 5 requests
            WindowStart:  time.Now(),         // start of 60s window
            WindowLength: 60 * time.Second,   // window duration
        }
        UserLimits[userID] = limit
    }

    // Reset counter if the 60-second window has passed
    if time.Since(limit.WindowStart) > limit.WindowLength {
        limit.Requests = 0
        limit.WindowStart = time.Now()
    }

    // If user has hit the max requests → block
    if limit.Requests >= limit.MaxRequests {
        return false
    }

    // Count this request and allow it
    limit.Requests++
    return true
}
