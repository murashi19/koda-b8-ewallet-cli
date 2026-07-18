FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ewallet .

FROM postgres:alpine

WORKDIR /app

COPY --from=builder /app/ewallet .

CMD ["./ewallet"]