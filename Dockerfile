FROM golang:1.22-alpine AS build

RUN apk add --no-cache curl ca-certificates

WORKDIR /app

RUN curl -sLo /usr/local/bin/tailwindcss https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 \
    && chmod +x /usr/local/bin/tailwindcss

RUN go install github.com/a-h/templ/cmd/templ@latest

COPY go.* ./

RUN go mod download && go mod verify

COPY . .

RUN templ generate && tailwindcss -i internal/view/css/styles.css -o static/css/styles.css

RUN CGO_ENABLED=0 GOOS=linux go build -o nit ./cmd/main.go

FROM scratch AS deploy

WORKDIR /app

COPY --from=build /app/static ./static

COPY --from=build /app/.env .

COPY --from=build /app/nit .

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 3000

ENTRYPOINT [ "/app/nit" ]
