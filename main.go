package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jintoples/gin-restful/app"
	"github.com/jintoples/gin-restful/controllers"
	"github.com/jintoples/gin-restful/repository"
	"github.com/jintoples/gin-restful/services"
)

func main() {
	db := app.NewDb()
	bookRepository := repository.NewBookRepository()
	bookService := services.NewBookServiceImpl(bookRepository, db)
	bookController := controllers.NewBookController(bookService)

	r := gin.Default()
	r.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/book", bookController.FindAll)
	r.GET("/book/:id", bookController.FindById)
	r.POST("/book", bookController.Create)
	r.PUT("/book/:id", bookController.Update)
	r.DELETE("/book/:id", bookController.Delete)

	r.Run(":3000")
}
