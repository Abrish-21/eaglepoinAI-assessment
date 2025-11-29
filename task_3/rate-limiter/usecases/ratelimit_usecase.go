package usecases

import (
    "time"
    "rate-limiter/entities"
)

// In-memory storage
var UserLimits = map[string]*entities.RateLimit{}

// Check if request is allowed
func Allow(userID string) bool {
    limit, ok := UserLimits[userID]
    if !ok {
        limit = &entities.RateLimit{
            Requests:     0,
            MaxRequests:  5,
            WindowStart:  time.Now(),
            WindowLength: 60 * time.Second,
        }
        UserLimits[userID] = limit
    }

    if time.Since(limit.WindowStart) > limit.WindowLength {
        limit.Requests = 0
        limit.WindowStart = time.Now()
    }

    if limit.Requests >= limit.MaxRequests {
        return false
    }

    limit.Requests++
    return true
}
