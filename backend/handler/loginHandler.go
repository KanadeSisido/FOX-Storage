package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h userHandler) LoginHandler(ctx *gin.Context) {

	var loginForm LoginForm
	err := ctx.ShouldBind(&loginForm)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid json",
		})
		return
	}

	c := ctx.Request.Context()
	token, err := h.controller.LoginController(c, loginForm.Username, loginForm.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.SetCookie("Authorization", *token, 3600, "/", "localhost", false, true)
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.JSON(http.StatusAccepted, gin.H{"message": "logined successfully"})
}
