# Eaglepoint AI Assessment

**Repository:** [eaglepointAI-assessment](#)

This repository contains solutions for three technical assessment tasks submitted to Eaglepoint AI.

---

## Task 1: Smart Text Analyzer

**Language:** Python

### Approach:

* Used Python for its **readability** and **strong built-in support for text processing**.
* Implemented an **OOP-style scalable solution** for future enhancements (e.g., sentiment analysis).
* Preprocessing includes:

  * Converting text to lowercase
  * Removing punctuation
* Word counting implemented using a **dictionary** to handle repeated words correctly.

### Notes:

* Designed to be **extensible** for adding new text analysis features without changing core logic.

---

## Task 2: Async Data Fetcher with Retry

**Language:** JavaScript

### Approach:

* Created a `fetchWithRetry` function that:

  * Fetches data from a URL
  * Retries on failure up to a maximum count
  * Waits **1 second** between retries using `async/await`
* Included a **mock API function** to simulate random success/failure for testing.
* Handled asynchronous fetch properly using **Promises and recursion**, avoiding multiple simultaneous retries.

### Key References:

* [MDN Fetch API](https://developer.mozilla.org/en-US/docs/Web/API/Window/fetch)
* [StackOverflow: Retry async function with delay](https://stackoverflow.com/questions/73073043/how-to-retry-an-async-function-with-a-delay-in-javascript)

---

## Task 3: Rate Limiter

**Language:** Go

**Architecture:** Clean Architecture (Uncle Bob's ) 

### Requirements:

* Limit: 5 requests per 60 seconds per user
* Track requests **per user ID**
* Block requests when limit is exceeded
* Auto-reset after the time window
* Include working examples

### Approach:

* Used **in-memory storage** (`map`) to track user requests.
* Created `RateLimitMiddleware` to enforce limits before passing requests to handlers.
* Simple `Allow` function in `usecases` checks requests and resets counters based on window time.
* No database or external repo needed for simplicity and interview-readiness.

### Folder Structure:

```
rate-limiter/
├── cmd/
│   └── main.go                # HTTP server + routes
├── entities/
│   └── ratelimit.go           # RateLimit struct
├── usecases/
│   └── ratelimit_usecase.go   # Logic for counting requests and checking limits
└── middleware/
    └── ratelimit_middleware.go # Enforces rate limit per request
```

### Running the Server:

```bash
cd rate-limiter
go run ./cmd/main.go
```

### Testing:

* Use **curl** to simulate requests:

```bash
curl -H "X-User-ID: alice" http://localhost:8080/
```

* Allowed requests: first 5 per 60 seconds
* Exceeded requests: returns `Rate limit exceeded` (HTTP 429)
* Reset automatically after 60 seconds
* Test with multiple users to verify per-user tracking

### Key References:

* [Go `http` package](https://pkg.go.dev/net/http#ListenAndServe)
* [curl documentation](https://curl.se/docs/manpage.html)

