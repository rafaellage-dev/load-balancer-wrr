FROM golang:1.21-alpine
WORKDIR /app
COPY main.go .
RUN go build -o backend-service main.go
EXPOSE 8080
CMD ["./backend-service"]