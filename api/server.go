package api

import (
	db "example/backend/db/sqlc"
	"github.com/gin-gonic/gin"
)

//Server serves HTTP requests for services
type Server struct {
	store  db.Store
	router *gin.Engine
}

func CreateServer(store *db.Store) *Server {
	router := gin.Default()
	server := &Server{store: *store}
	server.router = router
	return server
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: *store}
	router := gin.Default()

	router.POST("/person", server.createPerson)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	server.router.Run(address)
	return nil
}

// Handlers
func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
