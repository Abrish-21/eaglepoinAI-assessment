# Eaglepoint AI Assessment

You can find my unfiltered thought process here: [https://docs.google.com/document/d/1N0fOZS46YpvOvdr7M48pZg9axYyI7sX4BradfM1tlrA/edit?usp=sharing]

This repository contains solutions for three technical assessment tasks submitted to Eaglepoint AI. This README documents the **solution approach, thought process, and implementation steps** for each task.

---

## Getting Started

1. Clone the repository:

```bash
git clone git@github.com:username/eaglepointAI-assessment.git
cd eaglepointAI-assessment
```

2. Navigate to the respective task folder to run the code.

---

## Task 1: Smart Text Analyzer

**Language:** Python

### Thought Process & Steps:

1. **Problem Understanding:** Count words in a text while handling punctuation and case sensitivity.

2. **Language Choice:** Python – chosen for readability, built-in text processing libraries, and familiarity from DSA work.

3. **Design Approach:**

   * Used **Object-Oriented Programming (OOP)** to structure the solution for **scalability and maintainability** rather than a single procedural function.
   * This allows future enhancements without modifying existing logic.

     * For example, if we want to add **sentiment analysis**, we can simply add a `get_sentiment()` method to the class without changing the word counting or preprocessing logic.
   * Added a **preprocessing function** to:

     * Convert text to lowercase
     * Remove punctuation ([GeeksforGeeks reference](https://www.geeksforgeeks.org/python/python-remove-punctuation-from-string/))
   * Word counting is handled using a **dictionary** to correctly handle repeated words.

4. **Edge Cases:** Handled punctuation differences like `fox` vs `fox.` to ensure consistent word counting.

5. **Outcome:** Achieved accurate word count with a design that allows **future enhancements** like sentiment analysis or other text processing tasks.

---

## Task 2: Async Data Fetcher with Retry

**Language:** JavaScript

### Thought Process & Steps:

1. **Problem Understanding:** Fetch data from an API and retry on failure with a delay.

2. **Challenges:**

   * `fetch()` returns Promises, so using simple `try/catch` could miss asynchronous errors.
   * Needed to **wait 1 second between retries** to avoid overwhelming the server.

3. **Design Approach:**

   * Implemented `fetchWithRetry` using **async/await and recursion**.
   * Added a **mock API** that randomly succeeds or fails to simulate network unreliability.
   * Recursive approach ensures retries stop when a successful response is received or maximum attempts are reached.

4. **References:**

   * [MDN Fetch API](https://developer.mozilla.org/en-US/docs/Web/API/Window/fetch)
   * [StackOverflow: Retry async function with delay](https://stackoverflow.com/questions/73073043/how-to-retry-an-async-function-with-a-delay-in-javascript)

5. **Outcome:** Reliable fetch function with automatic retries and delay handling.

---

## Task 3: Rate Limiter

**Language:** Go
**Architecture:** Clean Architecture (simplified)

### Thought Process & Steps:

1. **Problem Understanding:** Limit API requests to 5 per 60 seconds per user, block excess requests, and auto-reset.

2. **Language Choice:** Go – chosen for speed, lightweight concurrency with goroutines, and ease of handling multiple requests efficiently.

3. **Architecture Choice:** Clean Architecture

   * Applied layers: **Entities → Interfaces → Usecase**
   * Focused on maintainability and scalability while keeping the implementation simple.
   * Removed the repository layer since in-memory storage suffices for this demonstration.

4. **Implementation Steps:**

   * **Entities:** Define `RateLimit` struct with request count and window information.
   * **Usecases:** Implement `Allow` function to check request limits and reset the window as needed.
   * **Middleware:** Implement `RateLimitMiddleware` to intercept HTTP requests and enforce limits.
   * **HTTP Server:** Set up routes and attach middleware.

5. **Testing & Debugging:**

   * Run the server:

     ```bash
     cd rate-limiter
     go run ./cmd/main.go
     ```
   * Test with `curl`:

     ```bash
     curl -H "X-User-ID: alice" http://localhost:8080/
     ```
   * Verify:

     * First 5 requests succeed
     * Requests above 5 within 60 seconds return HTTP 429
     * Different user IDs have separate limits
     * Rate limit resets after 60 seconds

6. **References:**

   * [Go `http` package](https://pkg.go.dev/net/http#ListenAndServe)
   * [curl documentation](https://curl.se/docs/manpage.html)
   * [Uncle Bob Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

### Folder Structure:

```
rate-limiter/
├── cmd/
│   └── main.go                # HTTP server + route setup
├── entities/
│   └── ratelimit.go           # RateLimit struct
├── usecases/
│   └── ratelimit_usecase.go   # Logic for counting requests and checking limits
└── middleware/
    └── ratelimit_middleware.go # Middleware to enforce rate limits
```

