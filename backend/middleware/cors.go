package middleware

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func CorsMiddleWare() gin.HandlerFunc {
	err := godotenv.Load()

	fmt.Println(err)

	conf := cors.DefaultConfig()
	fmt.Println(os.Getenv("FRONT_ORIGIN"))
	conf.AllowOrigins = []string{os.Getenv("FRONT_ORIGIN")}
	conf.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	conf.AllowCredentials = true
	return cors.New(conf)
}