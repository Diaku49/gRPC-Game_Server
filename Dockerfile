# Build Stage
FROM golang:1.25.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
 go -out game_server ./cmd/gamer_server

# Run Stage
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/game_server .

EXPOSE 50051

CMD ["./game_server"]