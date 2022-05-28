package server

import (
	"github.com/dedidot/grpc-book-server-streaming/model"
	"github.com/dedidot/grpc-book-server-streaming/proto"
	"github.com/dedidot/grpc-book-server-streaming/service"
	"log"
)

type Server struct {
	proto.BookServiceServer
}

func mapToBook(book model.Book) *proto.Book {
	return &proto.Book{
		Id:     book.Id,
		Title:  book.Title,
		Author: book.Author,
		IsRead: book.IsRead,
	}
}

func (s *Server) GetBookData(req *proto.GetBookRequest, stream proto.BookService_GetBookDataServer) error {
	log.Printf("Get book stream")
	books := service.GetBooks()
	for _, book := range books {
		data := mapToBook(book)
		_ = stream.Send(&proto.GetBookResponse{
			Status: true,
			Data:   data,
		})
	}

	return nil
}
