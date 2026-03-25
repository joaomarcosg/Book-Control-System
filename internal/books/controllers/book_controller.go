package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/joaomarcosg/Book-Control-System/internal/books/models"
)

type BooksController struct {
	bookService models.BookService
}

func NewBooksController(bookService models.BookService) *BooksController {
	return &BooksController{
		bookService: bookService,
	}
}

func (b *BooksController) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/books")

	{
		users.POST("", b.CreateBook)
		users.GET("/:id", b.GetBook)
		users.GET("", b.GetAllBooks)
		users.PUT("/:id", b.UpdateBook)
		users.DELETE("/:id", b.DeleteBook)
	}
}

func (b *BooksController) CreateBook(ctx *gin.Context) {
}

func (b *BooksController) GetBook(ctx *gin.Context) {
}

func (b *BooksController) GetAllBooks(ctx *gin.Context) {
}

func (b *BooksController) UpdateBook(ctx *gin.Context) {
}

func (b *BooksController) DeleteBook(ctx *gin.Context) {
}
