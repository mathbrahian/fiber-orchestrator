FROM golang:1.20-bullseye AS builder

EXPOSE 80

WORKDIR /usr/src

COPY . .

RUN go build -o fiber-orchestrator .

FROM debian:bullseye-slim

COPY --from=builder /usr/src/fiber-orchestrator /usr/local/bin/fiber-orchestrator

CMD ["fiber-orchestrator"]