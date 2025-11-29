package entities

import "time"
// This is a struct to hold rate limit info for a user
type RateLimit struct {
    Requests     int
    MaxRequests  int
    WindowStart  time.Time
    WindowLength time.Duration
}
