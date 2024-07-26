package service

import (
	"chat_controller_server/repository"
	"chat_controller_server/types/table"
)

type Service struct {
	repository *repository.Repository

	AvgServerList map[string]bool
}

func NewService(repository *repository.Repository) *Service {
	s := &Service{
		repository:    repository,
		AvgServerList: make(map[string]bool),
	}

	s.setServerInfo()

	return s
}

func (s *Service) GetAvailableServerList() ([]*table.ServerInfo, error) {
	return s.repository.GetAvailableServerList()
}

func (s *Service) setServerInfo() {
	if serverList, err := s.GetAvailableServerList(); err != nil {
		panic(err.Error())
	} else {
		for _, server := range serverList {
			s.AvgServerList[server.IP] = true
		}
	}
}
