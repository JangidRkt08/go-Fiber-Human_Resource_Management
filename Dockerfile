FROM golang:1.25.1-alpine  AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 3000
CMD ["./main"]