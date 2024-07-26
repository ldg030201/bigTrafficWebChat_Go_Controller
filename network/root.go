package network

import (
	"chat_controller_server/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engin *gin.Engine

	service *service.Service

	port          string
	avgServerList map[string]bool
}

func NewNetwork(service *service.Service, port string) *Server {
	s := &Server{
		engin:         gin.New(),
		service:       service,
		port:          port,
		avgServerList: make(map[string]bool),
	}

	s.engin.Use(gin.Logger())
	s.engin.Use(gin.Recovery())
	s.engin.Use(cors.New(cors.Config{
		AllowWebSockets:  true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	s.setServerInfo()

	return s
}

func (s *Server) Start() error {
	return s.engin.Run(s.port)
}

func (s *Server) setServerInfo() {
	if serverList, err := s.service.GetAvailableServerList(); err != nil {
		panic(err.Error())
	} else {
		for _, server := range serverList {
			s.avgServerList[server.IP] = true
		}
	}
}
