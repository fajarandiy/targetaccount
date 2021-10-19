package api

import (
	db "com.example/targetaccount/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP response for our services.
type Server struct {
	store  db.Store
	router *gin.Engine
}

//NewServer create a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/targetaccount", server.createTargetAccount)
	router.GET("/targetaccount/:id", server.getTargetAccount)
	router.GET("/targetaccount", server.getListTargetAccount)
	router.POST("/targetaccount/update", server.updateTargetAccount)
	router.POST("/targetaccount/delete", server.deleteTargetAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a spesific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
