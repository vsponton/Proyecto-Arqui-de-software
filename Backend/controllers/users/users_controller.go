package users

import (
	userService "backend/services/users"
	"net/http"

	"github.com/gin-gonic/gin"

	userDomain "backend/domain/users"
)


func Login(c *gin.Context){
	var userData userDomain.UserData
	c.BindJSON(&userData) 
	loginResponse := userService.Login(userData.User,userData.Password) //Llamo al servicio
	c.JSON(http.StatusOK, loginResponse) // muestro los datos
}