FROM golang:1.22.4

WORKDIR /app

COPY . .

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN go mod download
RUN go build -o app ./app/cmd/main.go

RUN chmod +x /app/scripts/wait_for_postgres.sh

EXPOSE 8081

CMD ["./app"]
