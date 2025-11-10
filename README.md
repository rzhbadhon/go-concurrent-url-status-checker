# Go Concurrent URL Status Checker (REST API)

<p align="center">
  <img src="https://miro.medium.com/v2/resize:fit:1358/format:webp/1*hlz9v7VdJDf7O-bCz14vIw.png" alt="Go Concurrent URL Status Checker Logo" width="150"/>
</p>

<p align="center">
  <b>A simple REST API written in Go that concurrently checks website statuses (Up/Down).</b>
</p>

---

## 📖 Overview

This project is a lightweight **Go web service** that demonstrates concurrency in action.  
It exposes an HTTP endpoint where you can POST a list of URLs, and it will check them **concurrently** to report whether they are reachable.

The project showcases Go’s concurrency primitives:

- 🌀 **Goroutines** for parallel execution  
- 📡 **Channels** for result collection  
- ⏳ **WaitGroup** for synchronization  

Perfect for learning Go’s concurrency model in a practical, API‑driven way.

---

## ⚙️ How It Works

1. A client sends a `POST` request to `/verify` with a JSON body containing URLs:
   ```json
   {
     "urls": [
       "https://google.com",
       "https://github.com",
       "https://invalid-url.test"
     ]
   }
2. The server spins up a goroutine for each URL.

3. Each goroutine performs an http.Get with a timeout.

4. Results are sent into a channel as models.Result objects.

5. Once all goroutines finish, the channel is closed.

6. The handler aggregates results into a JSON response:
7. 
```bash
  {
    "https://google.com": "Up ✅ (Status: 200 OK)",
    "https://github.com": "Up ✅ (Status: 200 OK)",
    "https://invalid-url.test": "Down 🔴 (Error: no such host)"
  }
```
## 🚀 Getting Started
### Prerequisites
  ~ Go 1.18+ installed on your system

### Run the server:
```bash
git clone https://github.com/rzhbadhon/go-concurrent-url-status-checker.git
cd go-concurrent-url-status-checker
go run main.go
```
The server will start on http://localhost:7080

## 📡 API Usage
### Endpoint
POST /verify

### Request Body:
```json
{
  "urls": ["https://example.com", "https://another.com"]
}
```
### Response
```json
{
  "https://example.com": "Up ✅ (Status: 200 OK)",
  "https://another.com": "Down 🔴 (Error: timeout)"
}
```
## 🛠 Project Structure

-- rest/api.go → Defines the /verify handler

-- utils/worker.go → Contains the CheckUrl function for URL checking

-- models/ → Request/response data structures

-- main.go → Entry point that starts the HTTP server

## 🎯 Key Features
~ Concurrent URL checking using goroutines

~ Timeout handling for slow/unresponsive sites

~ JSON API response for easy integration

~ Minimal dependencies (only Go standard library)
