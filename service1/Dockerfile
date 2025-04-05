# Use official Golang image as base
FROM golang:1.20-alpine as builder

WORKDIR /app

# Install git for go mod tidy
RUN apk add --no-cache git

# Copy go.mod and go.sum files first to cache dependencies
COPY go.mod go.sum ./
RUN go mod tidy || cat /root/.cache/go-build/log.txt

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Use a smaller image to run the app
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main ./

EXPOSE 3000

# Run the Go app
CMD ["./main"]
