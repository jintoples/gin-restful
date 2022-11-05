package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/jintoples/gin-restful/models/domain"
	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

func (repository *BookRepositoryImpl) FindAll(ctx *gin.Context, tx *gorm.DB) ([]domain.Book, error) {
	var books []domain.Book
	result := tx.Find(&books)

	return books, result.Error
}

func (repository *BookRepositoryImpl) FindById(ctx *gin.Context, tx *gorm.DB, bookId int) (domain.Book, error) {
	var book domain.Book
	result := tx.First(&book, bookId)

	return book, result.Error
}

func (repository *BookRepositoryImpl) Save(ctx *gin.Context, tx *gorm.DB, book domain.Book) (domain.Book, error) {
	result := tx.Create(&book)

	return book, result.Error
}

func (repository *BookRepositoryImpl) Update(ctx *gin.Context, tx *gorm.DB, book domain.Book) (domain.Book, error) {
	result := tx.Model(&book).Updates(book)

	return book, result.Error
}

func (repository *BookRepositoryImpl) Delete(ctx *gin.Context, tx *gorm.DB, book domain.Book) error {
	result := tx.Delete(&book)

	return result.Error
}
