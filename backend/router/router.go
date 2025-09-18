package router

import (
	"fox-storage/handler"
	"fox-storage/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(_handler *handler.Handler) *gin.Engine {

	router := gin.Default()
	router.Use(middleware.CorsMiddleWare())

	// Routes
	auth := router.Group("/auth")
	auth.POST("/login", _handler.LoginHandler)		// "/auth/login"
	auth.POST("/signup", _handler.SignUpHandler)	// "/auth/signup"

	perm := router.Group("/s")
	perm.Use(middleware.Auth)
	{
		perm.GET("/rootName", _handler.RootNameHandler)	// "/s/rootName"

		storage := perm.Group("/:folderId")

		storage.GET("/", _handler.FolderHandler)			// "/:folderId/" フォルダの内容を表示する
		
		file := storage.Group("/file")
		{
			file.POST("/upload", _handler.CreateFileHandler)	// "/:folderId/file/upload" ファイルアップロード
			file.GET("/:id", _handler.PingHandler)		// "/:folderId/file/:id" ファイルダウンロード
		}

		folder := storage.Group("/folder")
		{
			folder.POST("/create", _handler.CreateFolderHandler)	// "/:folderId/folder/create"
		}
		
	}


	return router
}

