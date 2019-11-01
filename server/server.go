package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Server holds the server handlers and dependencies
type Server struct {
	router *mux.Router
	logger *logrus.Logger
}

// New returns a new instance of Server initialised
// with a router and a logger.
func New(router *mux.Router, logger *logrus.Logger) *Server {
	return &Server{
		router: router,
		logger: logger,
	}
}

// InitRoutes registers all the routes for the server
func (s *Server) InitRoutes() {
	s.router.HandleFunc("/getexample", s.handleExampleGet)
	s.router.HandleFunc("/postexample", s.handleExamplePost)
}

// Run starts the server listening on the specified port
func (s *Server) Run(port string) {
	fmt.Printf("listening on port %v", port)
	log.Fatal(http.ListenAndServe(port, s.router))
}
