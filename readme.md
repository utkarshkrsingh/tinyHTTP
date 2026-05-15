<h2 align="center">TinyHTTP</h2>

A minimal HTTP/1.1 server built from scratch in Go without using the `net/http` package.  
The project demonstrates how HTTP works over raw TCP connections by manually handling request parsing and response generation.

---

## Features

- TCP server using Go’s `net` package  
- Manual HTTP request parsing  
- Header handling with `Content-Length` support  
- Structured request and response models  
- Valid HTTP/1.1 response generation  
- Method-based routing (`GET /`, `GET /hello`)  

---

## Architecture

```
TCP Connection → Request Parsing → Router → Handler → Response Writing
```

---

## Project Structure

```
/cmd/server → entry point
/internal/http → request, response, parser, writer
/internal/server → connection handling
/internal/handler → router and handlers
/internal/transport → connection utilities

```

---

## Run

```bash
go run cmd/server/main.go
```

Server runs on:
```
localhost:4221
```

---

## Example

```bash
curl http://localhost:4221/
```

Response:
```
HTTP/1.1 200 OK
Content-Length: 11
Connection: close

Hello World
```

---

Notes
- Built for learning and understanding HTTP internals
- Does not support keep-alive, HTTP/2, or advanced routing
