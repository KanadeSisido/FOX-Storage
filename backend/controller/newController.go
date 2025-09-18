package controller

import (
	"fox-storage/service"
)

type Controller struct {
	service *service.Service
}

func NewController(_service *service.Service) *Controller {
	_controller := Controller{service: _service}

	return &_controller
}