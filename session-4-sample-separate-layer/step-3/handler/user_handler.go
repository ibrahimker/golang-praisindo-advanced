package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ibrahimker/golang-praisindo-advanced/session-4-sample-separate-layer/step-3/service"
	"net/http"
)

// IUserHandler mendefinisikan interface untuk handler user
type IUserHandler interface {
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

// GetAllUsers menghandle permintaan untuk mendapatkan semua user
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users := h.userService.GetAllUsers()
	c.JSON(http.StatusOK, users)
}
