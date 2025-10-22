package router

import (
	"fox-storage/handler"
	"fox-storage/middleware"

	"github.com/gin-gonic/gin"
)

func NewItemsRouter(_itemsHandler handler.ItemsHandler, _userHandler handler.UserHandler) *gin.Engine {

	router := gin.Default()
	router.Use(middleware.CorsMiddleWare())

	// Routes
	auth := router.Group("/auth")
	auth.POST("/login", _userHandler.LoginHandler)   // "/auth/login"
	auth.POST("/signup", _userHandler.SignUpHandler) // "/auth/signup"

	perm := router.Group("/s")
	perm.Use(middleware.Auth)
	{
		perm.GET("/rootName", _userHandler.RootNameHandler) // "/s/rootName"

		storage := perm.Group("/:folderId")

		storage.GET("/", _itemsHandler.FileHandler) // "/:folderId/" フォルダの内容を表示する

		file := storage.Group("/file")
		{
			file.POST("/upload", _itemsHandler.CreateFileHandler) // "/:folderId/file/upload" ファイルアップロード
			file.GET("/:id", _itemsHandler.FileHandler)           // "/:folderId/file/:id" ファイルダウンロード
		}

		folder := storage.Group("/folder")
		{
			folder.POST("/create", _itemsHandler.CreateFolderHandler) // "/:folderId/folder/create"
		}

	}

	return router
}
