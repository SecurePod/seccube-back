FROM golang:1.21-alpine as builder

WORKDIR /app/

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine:latest

WORKDIR /app/
COPY --from=builder /app/main /app
RUN chmod +X main
CMD ["./main"]