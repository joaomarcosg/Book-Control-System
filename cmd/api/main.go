package main

import (
	"log"

	"github.com/gin-gonic/gin"
	bookController "github.com/joaomarcosg/Book-Control-System/internal/books/controllers"
	loanController "github.com/joaomarcosg/Book-Control-System/internal/loans/controllers"
	userController "github.com/joaomarcosg/Book-Control-System/internal/users/controllers"
)

func main() {
	router := gin.Default()

	userController := userController.NewUserController()
	bookController := bookController.NewBooksController()
	loanController := loanController.NewLoanController()

	userController.RegisterRoutes(router)
	bookController.RegisterRoutes(router)
	loanController.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
