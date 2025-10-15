# 🏠 Smart Home Device Management API (Golang RESTful Application)

## 📌 Overview
This project is a **RESTful API** built with Golang for managing smart home devices. It allows users to **register devices, authenticate, update status, and fetch real-time data** while utilizing Golang’s core concepts.

---

## ⏳ Project Duration: **8 Hours**
| Phase             | Tasks                                          | Duration |
|------------------|---------------------------------------------|----------|
| **Hour 1**      | Project setup, environment configuration     | 1 hour   |
| **Hour 2**      | Database setup (PostgreSQL, GORM)            | 1 hour   |
| **Hour 3**      | User authentication (JWT, bcrypt)            | 1 hour   |
| **Hour 4**      | Device model & CRUD API implementation       | 1 hour   |
| **Hour 5**      | Real-time updates (Goroutines & Channels)    | 1 hour   |
| **Hour 6**      | Security enhancements (SHA-256, CGO optional) | 1 hour   |
| **Hour 7**      | Testing (Unit & Integration tests)           | 1 hour   |
| **Hour 8**      | API documentation & deployment               | 1 hour   |

---

## 🛠 Tech Stack
- **Golang** (Core Language)
- **Gin** (Web framework)
- **GORM** (ORM for database)
- **PostgreSQL** (Database)
- **bcrypt & SHA256** (Security)
- **Goroutines & Channels** (Concurrency)
- **CGO (Optional)** (Interfacing with C/C++ for enhanced security)
- **JWT (JSON Web Token)** (Authentication)

---

---

## 🚀 Core Golang Concepts Used
1. **Goroutines & Channels** – Handle concurrent requests for real-time device updates.
2. **Pointers** – Optimize memory when passing device data between functions.
3. **Maps & Slices** – Store active devices and device metadata.
4. **Cryptography (bcrypt & SHA-256)** – Secure user passwords and device communication.
5. **Bitwise Operations** – Manage device states (ON/OFF, Active/Idle) efficiently.
6. **CGO (Optional)** – Use C-based security libraries for encryption.

---

## 📌 API Endpoints
| Method | Endpoint          | Description             | Auth |
|--------|------------------|------------------------|------|
| POST   | /register        | Register a user        | ❌  |
| POST   | /login           | Authenticate user      | ❌  |
| POST   | /devices         | Register a device      | ✅  |
| GET    | /devices         | List all devices       | ✅  |
| PUT    | /devices/:id     | Update device status   | ✅  |
| DELETE | /devices/:id     | Remove a device        | ✅  |

---















