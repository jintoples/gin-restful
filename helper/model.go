package helper

import (
	"github.com/jintoples/gin-restful/models/domain"
	"github.com/jintoples/gin-restful/models/web"
)

func ToBookResponse(book domain.Book) web.BookResponse {
	return web.BookResponse{
		ID:          book.ID,
		Name:        book.Name,
		Author:      book.Author,
		Publication: book.Publication,
	}
}

func ToBookResponses(books []domain.Book) []web.BookResponse {
	var BookResponses []web.BookResponse
	for _, book := range books {
		BookResponses = append(BookResponses, ToBookResponse(book))
	}
	return BookResponses
}
