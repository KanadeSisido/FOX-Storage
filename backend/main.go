package main

import (
	"fmt"
	"fox-storage/controller"
	"fox-storage/db"
	"fox-storage/handler"
	"fox-storage/repository"
	"fox-storage/router"
	"fox-storage/service"
)

func main() {

	fmt.Println("FOX-Storage Started")
	db, err := db.InitDB()

	if err != nil {
		fmt.Println(err)
		fmt.Println("DB initialize Error.")
		return
	}

	// Create Router
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	controller := controller.NewController(service)
	handler := handler.NewHandler(controller)
	router := router.NewRouter(handler)

	router.Run()
}