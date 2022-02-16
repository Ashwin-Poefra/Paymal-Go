package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/gyataro/paymal/db/sqlc"
)

// Server serves HTTP requests
type Server struct {
	store  db.Store
	router *gin.Engine
}

// Creates new HTTP server instance
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/api/accounts", server.createAccount)
	router.GET("/api/accounts/:id", server.getAccount)
	router.GET("/api/accounts", server.listAccounts)

	router.POST("/api/transfers", server.createTransfer)

	router.POST("/api/users", server.createUser)

	server.router = router
	return server
}

// Starts HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
