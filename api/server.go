package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/themeszone/gobank/db/sqlc"
)

// Server serves HTTP requests for out banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.PATCH("/account", server.updateAccount)
	router.DELETE("/account/:id", server.deleteAccount)

	// add routes to router
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(e error) gin.H {
	return gin.H{
		"error": e.Error(),
	}
}
