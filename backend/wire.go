//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"fox-storage/controller"
	"fox-storage/db"
	"fox-storage/handler"
	"fox-storage/repository"
	"fox-storage/router"
	"fox-storage/service"
)

func InitializeRouter() *gin.Engine {
	wire.Build(
		db.InitDB,
		repository.NewItemRepository,
		repository.NewUserRepository,
		service.NewItemService,
		service.NewUserService,
		controller.NewItemController,
		controller.NewUserController,
		handler.NewItemHandler,
		handler.NewUserHandler,
		router.NewRouter,
	)

	return nil
}
