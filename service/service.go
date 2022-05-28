package service

import (
	"github.com/dedidot/grpc-book-server-streaming/model"
	"github.com/dedidot/grpc-book-server-streaming/repository"
)

func GetBooks() []model.Book {
	return repository.GetBooks()
}
