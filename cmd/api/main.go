package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	bookController "github.com/joaomarcosg/Book-Control-System/internal/books/controllers"
	loanController "github.com/joaomarcosg/Book-Control-System/internal/loans/controllers"
	userController "github.com/joaomarcosg/Book-Control-System/internal/users/controllers"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("BOOK_CONTROL_SYSTEM_DATABASE_USER"),
		os.Getenv("BOOK_CONTROL_SYSTEM_DATABASE_PASSWORD"),
		os.Getenv("BOOL_CONTROL_SYSTEM_DATABASE_HOST"),
		os.Getenv("BOOK_CONTROL_SYSTEM_DATABASE_PORT"),
		os.Getenv("BOOK_CONTROL_SYSTEM_DATABASE_NAME"),
	))

	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

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
