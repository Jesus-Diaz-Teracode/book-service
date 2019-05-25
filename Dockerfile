FROM golang:stretch AS build

WORKDIR /go/src/github.com/Jesus-Diaz-Teracode/book-service

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o books_service .



FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=build /go/src/github.com/Jesus-Diaz-Teracode/book-service/books_service .

CMD ["./books_service"]
