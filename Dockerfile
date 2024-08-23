FROM golang:1.22-alpine AS builder

  RUN apk add --no-cache curl
  WORKDIR /app
  RUN curl -sLo /usr/local/bin/tailwindcss https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 && chmod +x /usr/local/bin/tailwindcss
  RUN go install github.com/a-h/templ/cmd/templ@latest
  RUN go install github.com/pressly/goose/v3/cmd/goose@latest
  COPY go.* ./
  RUN go mod download && go mod verify
  COPY . .
  RUN templ generate && tailwindcss -i internal/view/css/styles.css -o static/css/styles.css
  RUN CGO_ENABLED=0 GOOS=linux go build -o nit ./cmd/main.go
  
FROM alpine:3.20 AS deployer

  WORKDIR /app
  COPY --from=builder /app/migrations ./migrations
  COPY --from=builder /app/nit .
  COPY --from=builder /app/static ./static
  COPY --from=builder /app/.env .
  COPY --from=builder /go/bin/goose /usr/local/bin/goose
  RUN apk add --no-cache ca-certificates
  CMD ["/bin/sh", "-c", "goose up && /app/nit"]
