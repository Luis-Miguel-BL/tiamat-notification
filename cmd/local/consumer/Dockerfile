FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/local/consumer
EXPOSE 8080
CMD ["./main"]
