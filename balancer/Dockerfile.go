FROM golang:1.21-alpine
WORKDIR /app
COPY main.go .
RUN go build -o loadbalancer main.go
EXPOSE 8000
CMD ["./loadbalancer"]
