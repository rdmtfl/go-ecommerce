FROM golang:1.23.8-alpine AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o gatewayApp ./cmd/api

RUN chmod +x /app/gatewayApp

####

FROM alpine:3.21.3

RUN mkdir /app

COPY --from=builder /app/gatewayApp /app

CMD [ "/app/gatewayApp" ]