# booking-api

A RESTful Booking API built using Go, Gin, and PostgreSQL.  
This service allows users to create and manage bookings with proper validation and conflict handling.

---

## 🚀 Design Choices & Assumptions

### 🧩 Architecture
- Followed a **layered architecture**:
  - `handlers` → HTTP request handling
  - `services` → business logic
  - `repository` → database interaction
  - `models` → data structures
- This improves **scalability, maintainability, and testability**

---

### ⚙️ Tech Stack
- **Go** → Fast and efficient backend
- **Gin** → Lightweight web framework
- **PostgreSQL** → Reliable relational database
- **Docker (optional)** → Easy setup and portability

---

### ✅ Validations Implemented
- ❌ Booking in the past is not allowed
- ❌ Invalid time format is rejected
- ❌ Start time must be before end time
- ❌ Overlapping bookings are prevented
- ❌ Coach must exist before booking

---

### 📌 Assumptions
- Each booking is associated with a **coach**
- Time is stored in **standard format (ISO / UTC)**
- No authentication is implemented (can be added later)
- System handles **basic concurrency via DB constraints/logic**

---

## 🛠️ Setup & Run Locally

### 🔹 Prerequisites
Make sure you have:
- Go installed (>= 1.20)
- PostgreSQL installed
- Git installed

---

### 🔹 1. Clone the Repository
```bash
git clone https://github.com/faizannq/booking-api.git
cd booking-api
