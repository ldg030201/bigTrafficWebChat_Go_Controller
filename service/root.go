package service

import (
	"chat_controller_server/repository"
)

type Service struct {
	repository *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	s := &Service{repository: repository}

	return s
}
