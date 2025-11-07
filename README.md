# Go Concurrent URL Status Checker

<p align="center">
  <img src="https://miro.medium.com/v2/resize:fit:1358/format:webp/1*hlz9v7VdJDf7O-bCz14vIw.png" alt="Go Concurrent URL Status Checker Logo" width="150"/>
</p>

<p align="center">
  <b>A simple command-line tool written in Go that concurrently checks website statuses (Up/Down).</b>
</p>

---

## 📖 Overview

This project is a lightweight CLI tool built with **Go** to demonstrate concurrency in action.  
It checks a list of websites **concurrently** and reports whether they are reachable.

The project showcases the **Worker Pool** pattern using:

- 🌀 Goroutines  
- 📡 Channels  
- ⏳ WaitGroup  

Perfect for learning Go’s concurrency model in a practical, hands-on way.

---

## ⚙️ How It Works

1. The program starts with a hardcoded list of websites.  
2. It creates two channels:  
   - `jobs` → distributes tasks  
   - `results` → collects outcomes  
3. A fixed number of worker goroutines form the worker pool.  
4. The `main` goroutine sends all URLs into the `jobs` channel, then closes it.  
5. Workers:  
   - Receive a URL from `jobs`  
   - Send an `http.Get` request  
   - Push the result (success/failure) into `results`  
6. When `jobs` is closed, workers finish and signal completion via `wg.Done()`.  
7. A separate goroutine waits for all workers (`wg.Wait()`), then closes `results`.  
8. The `main` goroutine ranges over `results`, prints outcomes, and exits gracefully.  

---

## 🚀 Getting Started

### Prerequisites
- Go 1.18+ installed on your system

### Run the tool
```bash
git clone https://github.com/yourusername/go-concurrent-url-status-checker.git
cd go-concurrent-url-status-checker
go run .
