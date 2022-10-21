package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/enevarez1/go-exercise/db/sqlc"
)

//HTTP Request for our service
type Server struct {
	store db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.GET("/users/:id", server.getUser)
	router.PUT("/users/:id", server.updateUser)
	router.DELETE("/users/:id", server.deleteUser)


	server.router = router
	return server
}

func(server *Server) Start (address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}