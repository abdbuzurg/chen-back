package routes

import (
	"chen/pkg/controller"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func serverListen(r *gin.Engine) {
	var port, host string

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8080"
	}

	host, exists = os.LookupEnv("HOST")
	if !exists {
		host = "localhost"
	}

	err := r.Run(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		fmt.Println(err.Error())
		panic("PIZDEC")
	}
}

func RunRoutes() {
	r := gin.Default()

	r.Use(cors.Default())
	r.Use(gin.Recovery())

	attachRoutes(r)
	serverListen(r)
}

func attachRoutes(r *gin.Engine) {
	authEndpoints(r)
	organizationCRUDEndpoints(r)
}

func authEndpoints(r *gin.Engine) {
	auth := r.Group("/auth")
	auth.POST("/register", controller.Register)
	auth.GET("/login", controller.Login)
}

func organizationCRUDEndpoints(r *gin.Engine) {
	org := r.Group("/organization")
	org.GET("/:id", controller.OrgFind)
	org.POST("", controller.OrgCreate)
	org.PUT("/:id", controller.OrgUpdate)
	org.DELETE("/:id", controller.OrgDelete)
}
