package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ibrahimker/golang-praisindo-advanced/assignment-1-risk-assessment/service"
)

// IUserHandler mendefinisikan interface untuk handler user
type IUserHandler interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetAllUsers(c *gin.Context)
}

type UserHandler struct {
	userService service.IUserService
}

// NewUserHandler membuat instance baru dari UserHandler
func NewUserHandler(userService service.IUserService) IUserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *UserHandler) GetUser(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *UserHandler) GetAllUsers(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
