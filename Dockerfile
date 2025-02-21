# Stage 1: Build the application
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application binary
RUN CGO_ENABLED=0 GOOS=linux go build -o survey_backend ./main.go

# Stage 2: Build a lightweight image to run the application
FROM alpine:latest

# Install necessary runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Set timezone
ENV TZ=Asia/Kolkata

# Set the working directory
WORKDIR /app

# Copy the binary and environment file from the builder stage
COPY --from=builder /app/survey_backend .
COPY --from=builder /app/.env .

# Expose the port your application will run on
EXPOSE 3031

# Set the entrypoint command
CMD ["./survey_backend"]
