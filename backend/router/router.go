package router

import (
	"fox-storage/handler"
	"fox-storage/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(_itemHandler handler.ItemHandler, _userHandler handler.UserHandler) *gin.Engine {

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

		storage.GET("/", _itemHandler.FolderHandler) // "/:folderId/" フォルダの内容を表示する

		file := storage.Group("/file")
		{
			file.POST("/upload", _itemHandler.CreateFileHandler) // "/:folderId/file/upload" ファイルアップロード
			file.GET("/:id", _itemHandler.FileHandler)           // "/:folderId/file/:id" ファイルダウンロード
		}

		folder := storage.Group("/folder")
		{
			folder.POST("/create", _itemHandler.CreateFolderHandler) // "/:folderId/folder/create"
		}

	}

	return router
}
