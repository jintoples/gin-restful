package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/jintoples/gin-restful/models/domain"
	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll(ctx *gin.Context, tx *gorm.DB) ([]domain.Book, error)
	FindById(ctx *gin.Context, tx *gorm.DB, bookId int) (domain.Book, error)
	Save(ctx *gin.Context, tx *gorm.DB, book domain.Book) (domain.Book, error)
	Update(ctx *gin.Context, tx *gorm.DB, book domain.Book) (domain.Book, error)
	Delete(ctx *gin.Context, tx *gorm.DB, book domain.Book) error
}
