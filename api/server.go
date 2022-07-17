package api

import (
	db "github.com/freedommmoto/metamaskonline_backend/model/sqlc"
	"github.com/freedommmoto/metamaskonline_backend/tool"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config tool.ConfigObject
	store  *db.Queries
	router *gin.Engine
}

func NewServer(config tool.ConfigObject, store *db.Queries) (*Server, error) {

	router := gin.Default()
	server := &Server{
		config: config,
		store:  store,
		router: router,
	}

	router.GET("/user/id/:id", server.getUserByID)
	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
