FROM golang:1.23.6

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY src .

RUN go build -o app .

CMD ["./app"]

EXPOSE 8080
EXPOSE 3000