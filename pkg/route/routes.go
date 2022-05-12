package route

import (
	service "chen/pkg/service/global"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	return r
}

type ResponseFormatterFunc = service.ResponseFormatterFunc

func NewServer(db *gorm.DB) Server {
	r := gin.Default()

	r.Use(cors.Default())

	server := Server{
		DB:     db,
		Router: r,
	}

	return server
}

// Gets the ports and the host from env and runs the server based on those parameters
