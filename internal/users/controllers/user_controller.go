package controllers

import (
	"net/http"

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
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	id, err := c.userService.CreateUser(ctx, &user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, id)

}

func (c *UserController) GetUser(ctx *gin.Context) {

}

func (c *UserController) GetAllUsers(ctx *gin.Context) {

}

func (c *UserController) UpdateUser(ctx *gin.Context) {

}

func (c *UserController) DeleteUser(ctx *gin.Context) {

}
