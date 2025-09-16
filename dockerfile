FROM golang:1.22-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /out/api ./cmd/api

FROM alpine:3.20
RUN addgroup -S app && adduser -S app -G app
WORKDIR /app

RUN mkdir -p /app/data && chown -R app:app /app
COPY --from=builder /out/api /usr/local/bin/api

ENV ADDR=":8080" \
    DATA_DIR="/app/data"

EXPOSE 8080
USER app

HEALTHCHECK --interval=30s --timeout=3s --retries=3 CMD wget -qO- http://127.0.0.1:8080/api/health || exit 1

CMD ["api"]
