package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joaomarcosg/Book-Control-System/internal/users/controllers"
)

func main() {
	router := gin.Default()

	userController := controllers.NewUserController()
	userController.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
