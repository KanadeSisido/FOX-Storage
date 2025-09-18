package handler

import (
	"fox-storage/controller"
)

type Handler struct {
	controller *controller.Controller
}

func NewHandler(_controller *controller.Controller) *Handler {
	
	handler := Handler{controller: _controller}

	return &handler
}