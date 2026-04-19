<p align="center">
    <h2>TinyHTTP: A HTTP Server from Scratch</h2>
</p>

## Overview

TinyHTTP is a minimal HTTP/1.1 server built from scratch in Go using raw TCP sockets.

The goal of this project is to understand how HTTP works at the protocol level—without relying on high-level libraries.

---

## Features

* TCP server using Go's `net` package
* Manual HTTP request parsing
* Supports `GET /`
* Handles `Content-Length` for request bodies
* Generates valid HTTP/1.1 responses

---

## How It Works

```text
Client → TCP connection → Byte stream → HTTP parsing → Response generation
```

### Request Handling Flow

1. Accept TCP connection
2. Read incoming byte stream (handle partial reads)
3. Detect end of headers (`\r\n\r\n`)
4. Parse request line and headers
5. Read body using `Content-Length`
6. Generate HTTP response

---

## Example

### Request

```http
GET / HTTP/1.1
Host: localhost
```

### Response

```http
HTTP/1.1 200 OK
Content-Type: text/plain
Content-Length: 5

Hello
```

---

## How to Run

```bash
go run main.go
```

Then in another terminal:

```bash
curl http://localhost:4220
```

---

## Project Structure

```
.
├── go.mod
├── httpcore
│   ├── parser.go
│   ├── request.go
│   └── response.go
├── main.go
├── protocol
│   ├── crlf.go
│   ├── methods.go
│   └── status.go
└── server.go
```

---

## Limitations (Intentional)

* Only supports HTTP/1.1
* Only `GET /` is implemented
* No persistent connections (keep-alive)
* No chunked transfer encoding
* Minimal header validation

---

## Learning Goals

* Understand HTTP over TCP
* Handle streaming data correctly
* Implement protocol parsing manually
* Build a server without frameworks

---

## Future Improvements

* Add routing system
* Serve static files
* Support multiple methods
* Implement keep-alive connections

---

## Why This Project

Most developers use HTTP daily but never implement it.

This project focuses on:

* understanding protocol internals
* handling raw network streams
* building systems from first principles
