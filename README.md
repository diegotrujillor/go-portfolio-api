# go-portfolio-api

A Go backend API (Mid-level) built as a portfolio project: clean structure, maintainable foundations, and ready to evolve into auth, persistence, testing, and CI/CD.

## Tech Stack

- Go (`net/http`)
- Docker (multi-stage build)
- Makefile for common workflows

## Project Structure

```text
go-portfolio-api/
├── cmd/
│   └── api/
│       └── main.go          # application entrypoint
├── internal/
│   ├── handlers/            # HTTP handlers (transport layer)
│   ├── services/            # application logic (use cases)
│   └── domain/              # entities and business rules
├── configs/                 # configuration (e.g., .env.example)
├── docs/                    # documentation and notes
├── scripts/                 # helper scripts (migrations, etc.)
├── Makefile
├── Dockerfile
└── go.mod
```

## Why cmd/ + internal/?

- cmd/api makes it explicit which executable/binary you are building and running.
- internal/ enforces encapsulation: packages inside internal can only be imported within this repository. This helps prevent accidental coupling and keeps boundaries clean as the project grows.

## Requirements

- Go 1.22+ (or the Go version you decide to standardize on)
- Docker (optional but recommended)

## How to Run (Local)

```bash
go mod tidy
go run ./cmd/api
```

## Verify Healthcheck

```bash
curl http://localhost:8080/ping
# {"message":"pong"}
```

## How to Run (Docker)

```bash
docker build -t go-portfolio-api:local .
docker run --rm -p 8080:8080 go-portfolio-api:local
```

## Makefile Commands

```bash
make run          # run the app locally
make test         # run tests
make lint         # (optional) linting if you add it later
make docker-build # build the Docker image
make docker-run   # run the Docker container
```
