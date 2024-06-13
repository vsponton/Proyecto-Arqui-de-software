package app

import (
	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()

	router.Use(cors.Default())
}

func StartRouter() {
	MapUrls()

	log.Info("Starting server")

	router.Run(":8080")

}
