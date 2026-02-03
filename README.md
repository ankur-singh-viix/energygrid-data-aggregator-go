# EnergyGrid Data Aggregator (Go)

## Overview
This project is a Go-based client application that fetches real-time telemetry data for **500 solar inverters** from a legacy EnergyGrid API while strictly respecting API constraints such as rate limiting, batching, and security requirements.

The mock backend API simulates a real-world constrained service, and this client demonstrates robust API integration techniques.

---

## Problem Constraints
The EnergyGrid API enforces the following rules:

- **Rate Limit:** Maximum **1 request per second**
- **Batch Size:** Maximum **10 devices per request**
- **Security:** Each request must include a cryptographic signature  
  `MD5(URL + Token + Timestamp)`
- **Error Handling:** API returns `HTTP 429` when rate limit is exceeded

---

## Solution Approach

### 1. Serial Number Generation
- Programmatically generates **500 serial numbers**:  
  `SN-000` to `SN-499`

### 2. Batching Strategy
- Serial numbers are split into **50 batches**
- Each batch contains **10 devices**

### 3. Rate Limiting
- Uses Go’s `time.Ticker` to strictly enforce **1 request per second**
- Ensures no accidental rate-limit violations

### 4. Retry Handling
- Automatically retries when `HTTP 429 (Too Many Requests)` is received
- Prevents data loss while respecting API constraints

### 5. Aggregation
- Responses from all batches are aggregated into a single in-memory dataset
- Final result contains telemetry for all **500 devices**

---

## Project Structure

energygrid-client-go/
│
├── cmd/
│ └── main.go # Application entry point
│
├── internal/
│ ├── api/
│ │ └── client.go # API & signature logic
│ ├── limiter/
│ │ └── rate_limiter.go 
│ └── utils/
│ └── helpers.go 
│
├── go.mod
└── README.md


---

## Setup Instructions

### 1. Start the Mock API Server
The mock API simulates the EnergyGrid backend and is provided in Node.js.

```bash
cd mock-api
npm install
npm start


Expected output:
   EnergyGrid Mock API running on port 3000
   Constraints: 1 req/sec, Max 10 items/batch


2. Run the Go Client

Open a new terminal:
    cd energygrid-client-go
    go run cmd/main.go

Sample Output:
    EnergyGrid Data Aggregator (Go)
   Total Batches: 50
    Processing batch 1/50
    Rate limited. Retrying...
    ...
    Processing batch 50/50
    Aggregation Complete
    Total Devices Fetched: 500
