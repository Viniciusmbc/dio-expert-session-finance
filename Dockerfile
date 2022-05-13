FROM golang:latest

WORKDIR /app

COPY go.mod ./

Run go mod download

COPY . .

RUN CGO_ENABLED=0 go build -v -a -installsuffix cgo -o finance ./cmd/server

EXPOSE 8080

CMD ["./finance"]
