<!-- repo_name = manthan_audit_app -->

FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env . 

RUN CGO_ENABLED=0 GOOS=linux go build -o repo_name ./main.go

FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata

ENV TZ=Asia/Kolkata

WORKDIR /app

COPY --from=builder /app/repo_name .
COPY --from=builder /app/.env .
EXPOSE 8000

CMD ["./repo_name"]
