package services

import (
	"github.com/gin-gonic/gin"
	"github.com/jintoples/gin-restful/models/web"
)

type BookService interface {
	FindAll(ctx *gin.Context) []web.BookResponse
	FindById(ctx *gin.Context, bookId int) (web.BookResponse, error)
	Create(ctx *gin.Context, request web.BookCreateRequest) (web.BookResponse, error)
	Update(ctx *gin.Context, request web.BookUpdateRequest) (web.BookResponse, error)
	Delete(ctx *gin.Context, bookId int) error
}
