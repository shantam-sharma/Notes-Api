# Build Stage
FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o notes-api ./cmd

# Runtime Stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/notes-api .

EXPOSE 8080

CMD ["./notes-api"]

# Containers do not automatically inherit files
# from the host machine.

# Anything the container needs must be:

#1. Copied into the image OR

#2. Passed as environment variables OR

#3. Mounted as a volume
