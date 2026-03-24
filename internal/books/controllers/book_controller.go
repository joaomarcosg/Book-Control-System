package controllers

import "github.com/gin-gonic/gin"

type BooksController struct{}

func NewBooksController() *BooksController {
	return &BooksController{}
}

func (b *BooksController) RegisterRoutes(r *gin.RouterGroup) {
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
