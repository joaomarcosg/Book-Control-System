package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/joaomarcosg/Book-Control-System/internal/users/models"
)

type UserController struct {
	userService models.UserService
}

func NewUserController(userService models.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/users")

	{
		users.POST("", c.CreateUser)
		users.GET("/:id", c.GetUser)
		users.GET("", c.GetAllUsers)
		users.PUT("/:id", c.UpdateUser)
		users.DELETE("/:id", c.DeleteUser)
	}

}

func (c *UserController) CreateUser(ctx *gin.Context) {

}

func (c *UserController) GetUser(ctx *gin.Context) {

}

func (c *UserController) GetAllUsers(ctx *gin.Context) {

}

func (c *UserController) UpdateUser(ctx *gin.Context) {

}

func (c *UserController) DeleteUser(ctx *gin.Context) {

}
