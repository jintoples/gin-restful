package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jintoples/gin-restful/models/web"
	"github.com/jintoples/gin-restful/services"
)

type BookControllerImpl struct {
	BookService services.BookService
}

func NewBookController(bookService services.BookService) BookController {
	return &BookControllerImpl{
		BookService: bookService,
	}
}

func (controller *BookControllerImpl) FindAll(c *gin.Context) {
	bookResponses := controller.BookService.FindAll(c)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   bookResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *BookControllerImpl) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	bookResponse, err := controller.BookService.FindById(c, id)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
		}

		c.JSON(http.StatusNotFound, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   bookResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *BookControllerImpl) Create(c *gin.Context) {
	var request web.BookCreateRequest
	c.Bind(&request)

	bookResponse, err := controller.BookService.Create(c, request)
	if err != nil {
		panic(err)
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   bookResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *BookControllerImpl) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	var request web.BookUpdateRequest
	c.Bind(&request)

	request.ID = id

	bookResponse, err := controller.BookService.Update(c, request)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
		}

		c.JSON(http.StatusNotFound, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   bookResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *BookControllerImpl) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	err = controller.BookService.Delete(c, id)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
		}

		c.JSON(http.StatusNotFound, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	c.JSON(http.StatusOK, webResponse)
}
