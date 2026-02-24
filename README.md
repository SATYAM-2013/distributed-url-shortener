<p align="center">
  <h1 align="center">Distributed URL Shortener</h1>
  <p align="center">
    <strong>Production-Grade Distributed URL Shortening Service Built in Go</strong>
  </p>
  <p align="center">
    Designed & Developed by <strong>SATYAM SINHA</strong>
  </p>
</p>

---

## 🚀 Overview

Distributed URL Shortener is a production-ready backend system built using **Go**, designed with real-world distributed systems principles.

The service is:

- Stateless
- Horizontally scalable
- Redis-backed
- Dockerized
- Rate-limited
- Cache-optimized
- Cloud-deployment ready

This project demonstrates how large-scale backend systems are architected, deployed, and operated in production environments.

---

## 🛠 Tech Stack

| Layer | Technology |
|-------|------------|
| Language | Go (Golang) |
| Storage | Redis |
| Caching | In-memory LRU Cache |
| Rate Limiting | Redis-based sliding window |
| Containerization | Docker (Multi-stage build) |
| Production Image | Distroless |
| Frontend | Streamlit |
| Observability | Prometheus Metrics |

---

## 📸 Application Preview

<p align="center">
  <img src="screenshots/home.png" width="45%"/>
  <img src="screenshots/shortened.png" width="45%"/>
</p>

---

## ✨ Core Features

### 🔹 Stateless REST API
Designed for horizontal scaling behind a load balancer.

### 🔹 Redis-backed Persistence
Redis acts as the source of truth for URL mappings.

### 🔹 Distributed Rate Limiting
Prevents abuse using Redis-based sliding window counters.

### 🔹 In-Memory LRU Cache
Optimizes hot URL lookups for low-latency responses.

### 🔹 Production-Ready Containerization
- Multi-stage Docker builds
- Distroless runtime image
- Environment-driven configuration

### 🔹 Observability
- `/metrics` endpoint for Prometheus
- `/health` endpoint for health checks

---

## 🏗 System Architecture
Client (Streamlit UI / API Consumer)
↓
Go HTTP API (Stateless)
↓
Service Layer (Logic)
↓
LRU Cache
↓
Redis

### Architectural Principles

- Separation of concerns
- Clean layered design
- Stateless compute layer
- Distributed rate limiting
- Cloud-native configuration

---

## 📡 API Endpoints

| Method | Endpoint | Description |
|--------|----------|------------|
| POST | `/shorten` | Create a short URL |
| GET | `/{code}` | Redirect to original URL |
| GET | `/health` | Health check |
| GET | `/metrics` | Prometheus metrics |
| GET | `/api` | Service metadata |

---

🐳 Running with Docker (Recommended)

Start the backend and Redis automatically:

docker compose up --build

The API will be available at:

http://localhost:8080

Redis runs inside the Docker network and connects automatically to the backend service.

🖥 Running Locally (Manual Setup)

If you prefer running services manually without Docker Compose:

1️⃣ Start Redis
docker run -d -p 6379:6379 redis
2️⃣ Start Backend
go run cmd/api/main.go

Backend will start on:

http://localhost:8080
🎨 Running Frontend (Streamlit UI)

Start the frontend:

streamlit run app.py

Open in your browser:

http://localhost:8501
⚙ Environment Variables

The service supports the following environment variables:

Variable	Description	Default
PORT	API server port	8080
REDIS_URL	Redis connection string	127.0.0.1:6379
CACHE_SIZE	LRU cache size	100000
🔒 Rate Limiting Strategy

The system implements a Redis-backed sliding window rate limiter with:

Per API key enforcement

Configurable request limits

Distributed-safe implementation

Production-ready behavior

☁ Cloud Deployment Ready

This service is fully containerized and can be deployed to:

Render

Fly.io

AWS ECS / EC2

Google Cloud Platform

Azure

Kubernetes clusters

Key properties enabling cloud deployment:

Stateless instances

Environment-based configuration

Containerized runtime

No local disk dependency

🧠 Engineering Highlights

This project demonstrates:

Distributed systems design

Caching strategies

Horizontal scalability

Rate limiting implementation

Docker multi-stage builds

Production container hardening

Clean layered architecture

Observability integration

📈 Potential Enhancements

Custom domains

URL expiration

Click analytics dashboard

Authentication layer

Kubernetes deployment

CI/CD automation

Load testing benchmarks

📄 License

MIT License

<p align="center"> Built with engineering discipline and production mindset. </p> 
