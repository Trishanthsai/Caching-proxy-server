# 🚀 Caching Proxy Server (Go)

A lightweight in-memory caching reverse proxy built with Go.

This project acts as an intermediary between the client and the origin server. It caches HTTP responses in memory to reduce unnecessary network requests and improve response time.

## ✨ Features

- HTTP Server
- Reverse Proxy
- In-Memory Cache
- Cache HIT / MISS Detection
- Fast Response for Cached Requests

## 🛠️ Tech Stack

- Go
- net/http
- io
- Go Maps

## 📂 Current Progress

- [x] HTTP Server
- [x] HTTP Client
- [x] Reverse Proxy
- [x] In-Memory Cache
- [x] Cache HIT
- [x] Cache MISS
- [ ] TTL (Time-To-Live)
- [ ] Cache Invalidation
- [ ] Rate Limiting
- [ ] Concurrency (RWMutex)
- [ ] CLI Configuration
- [ ] Logging

## 📌 Learning Goals

This project is being built to understand:

- HTTP Internals
- Reverse Proxy
- Caching
- Go Networking
- Backend Infrastructure
- System Design Concepts

## 🚀 Next Milestone

- Implement TTL (Time-To-Live)
- Automatic Cache Expiration
- Production-ready Cache

---