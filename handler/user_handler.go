package handler

import (
	"clean_architecture_with_go/domain/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetUsers(c *gin.Context)
}

type userHandler struct {
	usecase.UserUsecase
}

func NewUserHandler(srv usecase.UserUsecase) UserHandler {
	return &userHandler{srv}
}

func (h *userHandler) GetUsers(c *gin.Context) {
	users, err := h.UserUsecase.GetUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, users)
}
