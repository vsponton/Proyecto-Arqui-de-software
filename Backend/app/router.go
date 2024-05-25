package router

import (
	"backend/controllers/users"

	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// MaxAge:           12 * time.Hour,
	}))
}

func MapUrls(engine *gin.Engine) {
	engine.POST("/users/login", users.Login)

}

func StartRoute() {
	MapUrls()

	log.Info("Starting server")
	err := router.Run(":8080")
	if err != nil {
		return
	} //8090

}
