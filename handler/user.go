package handler

import (
	"ardi_go/auth"
	"ardi_go/helper"
	"ardi_go/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUseHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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
	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		respon := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, respon)
		return
	}
	formatter := user.FormatUser(newUser, token)

	respon := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, respon)

}
func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		respon := helper.APIResponse("Login  failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, respon)
		return
	}
	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		respon := helper.APIResponse("Login  failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, respon)
		return
	}

	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		respon := helper.APIResponse("Login failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, respon)
		return
	}

	formatter := user.FormatUser(loggedinUser, token)
	respon := helper.APIResponse("Succesfuly Loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, respon)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		respon := helper.APIResponse("Email Checking  failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, respon)
		return
	}

	isEmailAvailible, err := h.userService.IsEmailAvailible(input)
	if err != nil {

		errorMessage := gin.H{"errors": "Server Error"}

		respon := helper.APIResponse("Email Checking  failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, respon)
		return
	}

	data := gin.H{
		"is_availble": isEmailAvailible,
	}
	metaMessage := "Email has been registered"
	if isEmailAvailible {
		metaMessage = "Email is Availible"
	}
	respon := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, respon)

}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	//dapat dari jwt
	userID := 1

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar succesfuly uploaded", http.StatusOK, "succes", data)

	c.JSON(http.StatusOK, response)
}
