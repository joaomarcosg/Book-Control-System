package controllers

import "github.com/gin-gonic/gin"

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/users")

	users.POST("", c.CreateUser)
}

func (c *UserController) CreateUser(ctx *gin.Context) {

}
