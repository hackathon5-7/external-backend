FROM golang:1.22.4

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o auth-app ./cmd/main.go

EXPOSE 8081

CMD ["./auth-app"]