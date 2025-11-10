package routes

import (
	"net/http"

	"example.com/rest-api/helper"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if helper.ContextErrors(err, context, http.StatusInternalServerError, "Value's could not be parsed") {
		return
	}

	err = user.Save()
	if helper.ContextErrors(err, context, http.StatusInternalServerError, "An error occured") {
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "New User has been created", "user": user})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if helper.ContextErrors(err, context, http.StatusInternalServerError, "Value's could not be parsed") {
		return
	}

	err = user.ValidateCredentials()
	if helper.ContextErrors(err, context, http.StatusInternalServerError, "Invalid credentials") {
		return
	}

	token, err := helper.GenerateTokens(user.Email, user.Id.String())
	if helper.ContextErrors(err, context, http.StatusInternalServerError, "Invalid Tokens") {
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})

}
