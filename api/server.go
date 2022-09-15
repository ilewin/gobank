package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/transparentideas/gobank/db/sqlc"
	"github.com/transparentideas/gobank/token"
	"github.com/transparentideas/gobank/util"
)

// Server serves HTTP requests for out banking service
type Server struct {
	config     *util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config *util.Config, store db.Store) (*Server, error) {

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymetricKey)
	if err != nil {
		return nil, err
	}

	server := &Server{config: config, store: store, tokenMaker: tokenMaker}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {

	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRouts := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRouts.POST("/accounts", server.createAccount)
	authRouts.GET("/account/:id", server.getAccount)
	authRouts.GET("/accounts", server.listAccounts)
	authRouts.PATCH("/account", server.updateAccount)
	authRouts.DELETE("/account/:id", server.deleteAccount)
	authRouts.POST("/transfers", server.createTransfer)

	// add routes to router
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(e error) gin.H {
	return gin.H{
		"error": e.Error(),
	}
}
