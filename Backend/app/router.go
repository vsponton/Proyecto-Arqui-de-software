package app

import (
<<<<<<< HEAD
	cors "github.com/gin-contrib/cors"
=======
>>>>>>> b8ca57bfe323424126dfecabcf2b9fa4fc94d9f5
	"github.com/gin-gonic/gin"
)

<<<<<<< HEAD
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
=======
func AllowCORS(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
	c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Next()
}
>>>>>>> b8ca57bfe323424126dfecabcf2b9fa4fc94d9f5
