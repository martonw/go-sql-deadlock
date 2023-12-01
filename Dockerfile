# Build test.go 
FROM golang:1.21

WORKDIR /go/src/app
COPY . .

RUN go build test.go

CMD ["sh", "-c", "while true; do ./test; done"]
