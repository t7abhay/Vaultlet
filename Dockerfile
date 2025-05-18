FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o vaultlet ./cmd/Vaultlet

FROM debian:bookworm-slim


WORKDIR /app


COPY --from=builder /app/vaultlet .


EXPOSE 9090

CMD ["./vaultlet"]
