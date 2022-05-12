package server

import (
	"chen/pkg/route"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type server struct {
	Router            *gin.Engine
	DB                *gorm.DB
	ResponseFormatter ResponseFormatterFunc
}

func NewServer(db *gorm.DB) *server {
	r := route.NewRouter()

	server := &server{
		Router:            r,
		DB:                db,
		ResponseFormatter: DefaultResponseFormatterFunc,
	}

	server.initializeRoutes()

	return server
}

func (s *server) SetResponseFormatter(responseFormatterFunc ResponseFormatterFunc) {
	s.ResponseFormatter = responseFormatterFunc
}

// Adds the routes to the router
// Registers the routes that are used in the application
func (s *server) initializeRoutes() {
	route.AuthEndpoints(s.Router, s.DB, s.ResponseFormatter)
	route.OrganizationCRUDEndpoints(s.Router, s.DB, s.ResponseFormatter)
	route.BranchCRUDEndpoints(s.Router, s.DB, s.ResponseFormatter)
	route.HallCRUDEndpoints(s.Router, s.DB, s.ResponseFormatter)
	route.TableCRUDEndpoints(s.Router, s.DB, s.ResponseFormatter)
}

// Gets the ports and the host from env and runs the server based on those parameters
func (s *server) Listen() {
	var port, host string

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8080"
	}

	host, exists = os.LookupEnv("HOST")
	if !exists {
		host = "localhost"
	}

	err := s.Router.Run(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		fmt.Println(err.Error())
		panic("PIZDEC")
	}
}
