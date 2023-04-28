package handler

import (
	"ardi_go/helper"
	"ardi_go/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUseHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}
func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		respon := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, respon)
		return
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		respon := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, respon)
		return
	}
	formatter := user.FormatUser(newUser, "tokentoken")
	//token, err:= h.jwtService. GenerateToken()

	respon := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, respon)

}
