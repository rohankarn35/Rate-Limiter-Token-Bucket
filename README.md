# Token Bucket Rate Limiter in Go

This is a high-concurrency **rate limiter** implemented in Go using the **Token Bucket design pattern**. It demonstrates how to throttle requests effectively and safely using goroutines, mutexes, and channels.

---

## 🧠 Theory

**Rate Limiting** is the practice of restricting the rate at which an operation can be performed (e.g., API calls). It protects systems from abuse, ensures fair usage, and prevents overload.

**Token Bucket Pattern**:

- A bucket holds a number of tokens.
- Tokens are added at a fixed rate.
- Each request consumes a token.
- If no tokens are available, the request is denied or delayed.

---

## 📁 Project Structure

```
rate-limiter-token-bucket/
├── go.mod
├── main.go                  # Test file with concurrency
└── limiter/
    └── token_bucket.go      # Token Bucket implementation
```

---

## 🛠️ Setup

```bash
# Clone the repo
git clone https://github.com/rohankarn35/Rate-Limiter-Token-Bucket.git
cd rate-limiter-token-bucket

# Initialize Go module
go mod tidy

# Run the test
go run main.go
```

---

## 🔍 Token Bucket Configuration (in main.go)

```go
capacity := 10                  // Max tokens
refillRate := 5                 // Tokens added each interval
refillInterval := time.Second   // Interval for refilling tokens
```

You can tune these values to test different traffic conditions.

---

## 🧪 Sample Behavior

- Sends 50 concurrent requests.
- Spreads them over time with `time.Sleep`.
- Some requests are allowed (✅), others are rate limited (❌), depending on token availability.

---

## 🔐 Core Features

- Thread-safe token handling with `sync.Mutex`
- Background goroutine for token refill using `time.Ticker`
- Graceful shutdown of refill loop via `chan struct{}`
- Clean and modular structure

---

## 🚀 Future Enhancements

- Per-client rate limiting
- Middleware integration (HTTP/GRPC)
- Unit tests and benchmarking
- Redis-backed shared limiter (for distributed systems)

---

## 📄 License

Apache License

---

## 🤝 Author

**Rohan Karn**

---

Happy hacking! 🚀
