package service

import (
	"context"
	svc "github.com/Jesus-Diaz-Teracode/book-service/grpc"
	"log"
)

// GetBook returns a book
func (s *Server) GetBook(ctx context.Context, in *svc.BookRequest) (*svc.BookResponse, error) {
	book := &svc.BookResponse{
		Uuid: "ed03d133-4a24-433a-997c-a3bb22cfe869",
		Name: "Concurrency in Go",
		Isbn: 9781234567897,
	}

	log.Printf("Received request books store")

	return book, nil
}
