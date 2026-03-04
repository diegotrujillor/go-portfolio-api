# ---- build stage ----
FROM golang:1.26.0-alpine3.23 AS builder

WORKDIR /app

RUN apk add --no-cache ca-certificates git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o api ./cmd/api

# ---- runtime stage ----
FROM gcr.io/distroless/base-debian12:nonroot

# Copy certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy binary
COPY --from=builder /app/api /api

EXPOSE 8080

ENTRYPOINT ["/api"]