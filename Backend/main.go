package main

import (
	"cursos-ucc/app"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	app.MapUrls(engine)
	engine.Run(":8080")
}
