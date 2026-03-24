package main

import (
	"log"

	"github.com/gin-gonic/gin"
	bookController "github.com/joaomarcosg/Book-Control-System/internal/books/controllers"
	userController "github.com/joaomarcosg/Book-Control-System/internal/users/controllers"
)

func main() {
	router := gin.Default()

	userController := userController.NewUserController()
	bookController := bookController.NewBooksController()

	userController.RegisterRoutes(router)
	bookController.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
