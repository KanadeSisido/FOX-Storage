package service

import (
	"fox-storage/repository"
)

type Service struct {
	repository repository.Repository
}

func NewService(_repository *repository.Repository) *Service {

	service := Service{repository: *_repository}

	return &service
}