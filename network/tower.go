package network

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Tower struct {
	server *Server
}

func registerTowerAPI(server *Server) {
	t := &Tower{server: server}

	t.server.engin.GET("/server-list", t.serverList)
}

func (t *Tower) serverList(c *gin.Context) {
	response(c, http.StatusOK, t.server.service.AvgServerList)
}
