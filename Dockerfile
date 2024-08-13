FROM golang:1.22-alpine

RUN apk add --no-cache curl

WORKDIR /app

RUN curl -sLo /usr/local/bin/tailwindcss https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 \
    && chmod +x /usr/local/bin/tailwindcss

RUN go install github.com/a-h/templ/cmd/templ@latest

RUN go install github.com/air-verse/air@latest

COPY go.* ./

RUN go mod download && go mod verify

COPY . .

EXPOSE 42069

CMD ["air", "-c", ".air.toml"]
