FROM golang:1.23.8-alpine AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o orderApp ./cmd/api

RUN chmod +x /app/orderApp

####

FROM alpine:3.21.3

RUN mkdir /app

COPY --from=builder /app/orderApp /app

CMD [ "/app/orderApp" ]