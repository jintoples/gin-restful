package services

import (
	"github.com/gin-gonic/gin"
	"github.com/jintoples/gin-restful/helper"
	"github.com/jintoples/gin-restful/models/domain"
	"github.com/jintoples/gin-restful/models/web"
	"github.com/jintoples/gin-restful/repository"
	"gorm.io/gorm"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
	DB             *gorm.DB
}

func NewBookServiceImpl(bookRepository repository.BookRepository, DB *gorm.DB) BookService {
	return &BookServiceImpl{
		BookRepository: bookRepository,
		DB:             DB,
	}
}

func (service *BookServiceImpl) FindAll(ctx *gin.Context) []web.BookResponse {
	tx := service.DB.Begin()

	books, err := service.BookRepository.FindAll(ctx, tx)
	if err != nil {
		tx.Rollback()
	}

	return helper.ToBookResponses(books)
}

func (service *BookServiceImpl) FindById(ctx *gin.Context, bookId int) (web.BookResponse, error) {
	tx := service.DB.Begin()

	book, err := service.BookRepository.FindById(ctx, tx, bookId)
	if err != nil {
		tx.Rollback()
	}

	return helper.ToBookResponse(book), err
}

func (service *BookServiceImpl) Create(ctx *gin.Context, request web.BookCreateRequest) (web.BookResponse, error) {
	tx := service.DB.Begin()

	data := domain.Book{
		Name:        request.Name,
		Author:      request.Author,
		Publication: request.Publication,
	}

	book, err := service.BookRepository.Save(ctx, tx, data)
	if err != nil {
		tx.Rollback()
	}

	return helper.ToBookResponse(book), err
}

func (service *BookServiceImpl) Update(ctx *gin.Context, request web.BookUpdateRequest) (web.BookResponse, error) {
	tx := service.DB.Begin()

	_, err := service.BookRepository.FindById(ctx, tx, request.ID)
	if err != nil {
		tx.Rollback()
	}

	data := domain.Book{
		ID:          request.ID,
		Name:        request.Name,
		Author:      request.Author,
		Publication: request.Publication,
	}

	book, err := service.BookRepository.Update(ctx, tx, data)
	if err != nil {
		tx.Rollback()
	}

	return helper.ToBookResponse(book), err
}

func (service *BookServiceImpl) Delete(ctx *gin.Context, bookId int) error {
	tx := service.DB.Begin()

	book, err := service.BookRepository.FindById(ctx, tx, bookId)
	if err != nil {
		tx.Rollback()
	}

	err = service.BookRepository.Delete(ctx, tx, book)

	return err
}
