# Stage 1: Build the Go application
FROM golang:1.20-alpine as builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Copy the Go module files and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Create the final container
# FROM gcr.io/distroless/static
FROM alpine

WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/main .

# Copy the .env file
COPY .env .

# Expose the desired port
EXPOSE 8080

# Run the Go application
CMD ["./main"]
