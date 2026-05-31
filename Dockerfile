FROM golang:1.25

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o notes-api ./cmd

EXPOSE 8080

CMD ["./notes-api"]


# Containers do not automatically inherit files
# from the host machine.

# Anything the container needs must be:

#1. Copied into the image OR

#2. Passed as environment variables OR

#3. Mounted as a volume
