package app
import (

	"backend/controllers/users"
	"github.com/gin-gonic/gin"

)

func MapRoutes(engine *gin.Engine) { //levanta la aplicacion
	engine.POST("/users/login", users.Login)
	
}