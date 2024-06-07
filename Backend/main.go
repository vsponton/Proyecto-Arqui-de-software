package main

import (
	"cursos-ucc/app"
	"cursos-ucc/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.StartDbEngine()
	engine := gin.New()
	app.MapUrls(engine)
	engine.Run(":8080")
}
