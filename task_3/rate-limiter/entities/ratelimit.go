package entities

import "time"

type RateLimit struct {
    Requests     int
    MaxRequests  int
    WindowStart  time.Time
    WindowLength time.Duration
}
