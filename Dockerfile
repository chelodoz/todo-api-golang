# Build stage
FROM golang:1.18.4-alpine3.16 AS builder
WORKDIR /build

# Copy module files first so that they don't need to be downloaded again if no change
COPY go.* ./
RUN go mod download
RUN go mod verify

# Copy source files and build the binary
COPY . .
RUN go build -o todo ./cmd/todo

# Run stage
FROM alpine:3.16
WORKDIR /build
COPY --from=builder /build/todo .
COPY --from=builder /build/config/app.env .

EXPOSE 8080

CMD ["/build/todo"]