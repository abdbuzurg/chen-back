package route

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func NewServer(db *gorm.DB) Server {
	r := gin.Default()

	r.Use(cors.Default())
	r.Use(gin.Recovery())

	server := Server{
		DB:     db,
		Router: r,
	}

	server.registerRoutes()

	return server
}

func (s *Server) registerRoutes() {
	s.AuthEndpoints()
	s.OrganizationCRUDEndpoints()
	s.BranchCRUDEndpoints()
	s.HallCRUDEndpoints()
	s.TableCRUDEndpoint()
}

func (s *Server) ServerListen() {
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
