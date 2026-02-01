# Distributed URL Shortener

A production-grade, distributed URL shortener service built in Go.  
Designed with scalability, fault tolerance, and real-world backend engineering principles in mind.

This project demonstrates how large-scale backend services are architected, deployed, and operated in industry.

---

## 🚀 Key Features

- **Stateless REST API** enabling horizontal scaling
- **Redis-backed persistence** for URL mappings
- **Redis-based distributed rate limiting** (per API key)
- **In-memory LRU cache** for low-latency reads
- **Dockerized using multi-stage builds and distroless images**
- **Environment-based configuration** for cloud deployment
- **Health checks and observability hooks**
- **Production-ready error handling and validation**

---

## 🏗️ Architecture Overview

- **API Layer**: Go net/http handlers
- **Service Layer**: Business logic and abstractions
- **Cache Layer**: LRU cache for hot URLs
- **Storage Layer**: Redis as the source of truth
- **Rate Limiting**: Redis-backed sliding window counters
- **Deployment**: Docker + cloud hosting (Render/Fly/AWS-ready)

All application instances are stateless and can scale horizontally behind a load balancer.

---

## 🔒 Rate Limiting

Each request is protected using a Redis-backed rate limiter to prevent abuse and ensure fair usage across clients.

---

## ☁️ Deployment

The service is fully containerized and deployed to the cloud using Docker.  
Configuration is provided through environment variables, making the service portable across environments.

---

## 🧠 Why This Project

This project was built to practice and demonstrate **real-world backend system design**, including:
- Distributed systems
- Caching strategies
- Rate limiting
- Stateless services
- Cloud deployment workflows

It reflects patterns commonly used in large-scale production systems.
