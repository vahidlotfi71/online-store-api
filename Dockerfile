FROM golang:1.24-alpine AS builder

WORKDIR /app

# تنظیمات شبکه Go
ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=off

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine:latest

WORKDIR /app

# نصب CA certificates برای HTTPS
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]