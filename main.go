package main

import (
	svc "github.com/jedi4z/books-service/grpc"
	"github.com/jedi4z/books-service/service"
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
	svc.RegisterBookServer(s, &service.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}