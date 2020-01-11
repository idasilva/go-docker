FROM golang:alpine as builder
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]