# 🚦 Token Bucket Rate Limiter in Go

This project is a **high-concurrency rate limiter** implemented in Go using the **Token Bucket design pattern**. It demonstrates safe request throttling with goroutines, mutexes, and channels.

---

## 🧠 Theory

**Rate Limiting** controls how often an action can occur in a given timeframe, helping protect APIs from abuse and overload.

**Token Bucket Pattern**:

- A bucket holds tokens.
- Tokens refill at a fixed rate.
- Each request consumes one token.
- If no tokens are available, the request is denied or delayed.

---

## 📁 Current Project Structure

```
.
├── cmd
│   └── main.go
├── go.mod
├── internal
│   └── middleware
│       └── http.go
├── LICENSE
├── Makefile
├── pkg
│   └── limiter
│       ├── leaky_bucket.go
│       ├── limiter.go
│       └── token_bucket.go
├── README.md
└── tests
    ├── benchmark_test.go
    ├── concurrency_test.go
    └── limiter_test.go
```

---

## ⚡ Current Features

- **Token Bucket** and **Leaky Bucket** implementations
- Thread-safe with `sync.Mutex`
- Middleware-ready for HTTP
- Unit tests and benchmarks
- Modular Go project structure

---

## 🛠️ Setup

```bash
go mod tidy
go run ./cmd/main.go
```

---

## 🔮 Upcoming Advancements (Version 2)

We are evolving this into a **production-grade** rate limiter:

- **Per-IP** and **Per-API key** limiting
- Multiple algorithms: Token Bucket, Leaky Bucket, Sliding Window
- In-memory + Redis-backed storage for distributed deployments
- HTTP & gRPC middleware
- Configurable tiers per route or plan
- Prometheus metrics for monitoring
- Dockerized for easy deployment
- CI/CD with GitHub Actions

**Future Structure:**

```
.
├── cmd/server/
├── config/
├── deployments/
├── internal/api/
├── internal/server/
├── pkg/limiter/
├── pkg/storage/
├── test/
├── .github/workflows/
├── go.mod
├── LICENSE
├── Makefile
└── README.md
```

---

## 📄 License

Apache License 2.0

---

## 👨‍💻 Author

**Rohan Karn** — Backend & Go Developer
