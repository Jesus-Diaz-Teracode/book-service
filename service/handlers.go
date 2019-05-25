package service

import (
	"context"
	"github.com/jedi4z/gRPC-example/service-definitions/book-service"
	"log"
)

// GetBook returns a book
func (s *Server) GetBook(ctx context.Context, in *book_service.BookRequest) (*book_service.BookResponse, error) {
	book := &book_service.BookResponse{
		Uuid: "ed03d133-4a24-433a-997c-a3bb22cfe869",
		Name: "Concurrency in Go",
		Isbn: 9781234567897,
	}

	log.Printf("Received request books store")

	return book, nil
}
