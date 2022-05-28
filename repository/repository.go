package repository

import "github.com/dedidot/grpc-book-server-streaming/model"

func GetBooks() []model.Book {
	return []model.Book{
		{
			Id:     "satu",
			Title:  "Make mac",
			Author: "Micheal",
			IsRead: true,
		},
		{
			Id:     "dua",
			Title:  "Creating website",
			Author: "Peter",
			IsRead: true,
		},
	}
}
