<p align="center">
  <h1 align="center">🔗 Distributed URL Shortener</h1>
  <h3 align="center">Designed & Engineered by SATYAM SINHA</h3>
  <p align="center">
    Production-Grade Distributed URL Shortening Service Built in Go
  </p>
  <p align="center">
    Designed for Scalability • Fault Tolerance • Cloud Deployment
  </p>
</p>

A production-grade, horizontally scalable URL shortening service built in **Go**, designed using real-world backend engineering principles.

This system demonstrates distributed architecture patterns used in large-scale production services, including caching strategies, rate limiting, stateless APIs, and containerized cloud deployment.

---

##  System Architecture

```
                ┌─────────────────────┐
                │    Streamlit UI     │
                └──────────┬──────────┘
                           │
                           ▼
                ┌─────────────────────┐
                │  Go HTTP API Layer  │
                │  (Stateless)        │
                └──────────┬──────────┘
                           │
        ┌──────────────────┴──────────────────┐
        ▼                                     ▼
┌─────────────────┐                   ┌─────────────────┐
│   LRU Cache     │                   │     Redis       │
│ (Hot URL cache) │                   │  Source of Truth│
└─────────────────┘                   └─────────────────┘
                                              │
                                              ▼
                                      Rate Limiter
                                (Distributed via Redis)
```

---

##  Core Features

### ✅ Stateless REST API

* Designed for horizontal scaling
* Instances can be deployed behind a load balancer
* No in-memory state dependency

### ✅ Redis-backed Persistence

* URL mappings stored in Redis
* Ensures durability and distributed consistency

### ✅ Distributed Rate Limiting

* Redis-backed sliding window algorithm
* Prevents abuse per API key
* Cloud-safe across multiple instances

### ✅ In-Memory LRU Cache

* Reduces Redis read pressure
* Optimizes hot URL access
* Improves latency under load

### ✅ Dockerized (Production Ready)

* Multi-stage builds
* Distroless runtime image
* Minimal attack surface
* Portable across environments

### ✅ Environment-Based Configuration

* No hardcoded secrets
* Cloud-native configuration pattern
* Ready for Render / Fly.io / AWS / GCP

### ✅ Observability

* Prometheus metrics endpoint
* Health checks
* Structured logging

---

##  Design Principles

This project follows industry-standard backend architecture principles:

* Separation of concerns (API / Service / Storage layers)
* Stateless service design
* Cache-aside strategy
* Distributed rate limiting
* Container-first deployment
* Configuration via environment variables

These patterns mirror those used in high-scale systems at companies like Google, Meta, Amazon, and Netflix.

---

##  API Endpoints

### Create Short URL

```http
POST /shorten
```

Request:

```json
{
  "url": "https://example.com"
}
```

Response:

```json
{
  "short_url": "http://localhost:8080/abc123"
}
```

---

### Redirect

```http
GET /{code}
```

Redirects to the original URL.

---

### Health Check

```http
GET /health
```

---

### Metrics

```http
GET /metrics
```

Prometheus-compatible metrics.

---

##  Rate Limiting

* Implemented using Redis sliding window counters
* Configurable via environment variables
* Enforced per API key
* Distributed-safe across multiple instances

---

##  Environment Variables

| Variable            | Description            | Default        |
| ------------------- | ---------------------- | -------------- |
| PORT                | HTTP server port       | 8080           |
| REDIS_ADDR          | Redis address          | localhost:6379 |
| API_KEY             | API authentication key | required       |
| ENABLE_RATE_LIMIT   | Enable rate limiting   | true           |
| RATE_LIMIT          | Requests per window    | 100            |
| RATE_WINDOW_SECONDS | Time window            | 60             |
| CACHE_SIZE          | LRU cache size         | 100000         |

---

##  Docker Setup

### Start Redis

```bash
docker run -d -p 6379:6379 redis
```

### Run Application

```bash
docker build -t url-shortener .
docker run -p 8080:8080 url-shortener
```

---

##  Local Development

### 1️⃣ Start Redis

```bash
docker run -d -p 6379:6379 redis
```

### 2️⃣ Run Backend

```bash
go run cmd/api/main.go
```

### 3️⃣ Run Frontend (Streamlit)

```bash
streamlit run app.py
```

---

##  Deployment

The service is fully containerized and can be deployed to:

* Render
* Fly.io
* AWS ECS
* GCP Cloud Run
* Kubernetes clusters

Because the application is stateless, it supports horizontal scaling behind a load balancer without session affinity.

---

##  Scalability Strategy

* Stateless instances
* Redis as centralized storage
* Cache layer reduces database load
* Rate limiting distributed via Redis
* Docker enables portable scaling

Under load, multiple API instances can operate concurrently with shared Redis coordination.

---

##  Production Considerations

Potential enhancements for real-world deployment:

* Redis clustering
* URL expiration support
* Click analytics
* Persistent storage (RDB/AOF)
* Custom domains
* Authentication system
* OpenTelemetry tracing
* CI/CD pipeline integration

---

##  Tech Stack

* **Go**
* **Redis**
* **Docker**
* **Streamlit (Frontend UI)**
* **Prometheus (Metrics)**

---

##  License

MIT License

---
##  Connect With Me

LinkedIn: https://www.linkedin.com/in/satyam-sinha-b1646435b/
GitHub: https://github.com/satyam-2013
