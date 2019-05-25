package main

import (
	"github.com/jedi4z/gRPC-example/books-service/service"
	"github.com/jedi4z/gRPC-example/service-definitions/book-service"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

func main(){
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	book_service.RegisterBookServer(s, &service.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}