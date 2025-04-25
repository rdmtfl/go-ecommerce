FROM golang:1.18-alpine AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o gatewayApp ./cmd/api

RUN chmod +x /app/gatewayApp

####

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/gatewayApp /app

CMD [ "/app/gatewayApp" ]