# I know this is not the ideal way to write this dockerfile but whatever
FROM golang:1.25-alpine

LABEL org.opencontainers.image.source https://github.com/phl-code-club/coding-challenges

WORKDIR /var/app

COPY . .

RUN go run ./cmd/seed/

RUN go build -o coding-challenges .

ENTRYPOINT ["./coding-challenges"]
