FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download || { echo "Failed to download modules"; exit 1; }

COPY . .


RUN go build -o vibex-api cmd/api/main.go
RUN ls -l /app


FROM ubuntu:22.04

WORKDIR /root/

COPY --from=builder /app/vibex-api .


EXPOSE 8080
RUN chmod +x vibex-api

CMD ["./vibex-api"]
