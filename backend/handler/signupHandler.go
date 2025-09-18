package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignUpForm struct {
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) SignUpHandler(ctx *gin.Context) {
	
	var signUpForm SignUpForm
	err := ctx.ShouldBind(&signUpForm)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid type",
		})
		return
	}

	c := ctx.Request.Context()
	err = h.controller.SignupController(&c, signUpForm.Username, signUpForm.Email, signUpForm.Password)

	// TODO: 詳細にする
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
		return
		
	}
	ctx.JSON(http.StatusAccepted, gin.H{"message": "registered successfully"})
}