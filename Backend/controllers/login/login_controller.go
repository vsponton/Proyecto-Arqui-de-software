package loginController

import (
	"cursos-ucc/dto"
	"cursos-ucc/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Login(c *gin.Context) {
	var loginDto dto.LoginRequest
	err := c.BindJSON(&loginDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid request: %s", err.Error()),
		})
		return
	}
	log.Debug(loginDto)

	/*
		if err != nil {
			log.Error(err.Error())
			c.JSON(http.StatusBadRequest, dto.LoginResponse{
				"message": fmt.Sprintf("Invalid request: %s", err.Error()),
				//Message: fmt.Sprintf("Invalid request: %s", err.Error()),
				//Message: "Unauthorized login: Invalid credentials",
			})
			return
		}
		log.Debug(loginDto)
	*/

	var loginResponseDto dto.LoginResponse
	loginResponseDto, err = services.UserClient.Login(loginDto)

	token, err := services.Users(loginDto.Email, loginDto.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, loginResponseDto{
		Token: token,
	})
}

func Register(c *gin.Context) {
	var registerRequest dto.RegisterRequest
	err := c.BindJSON(&registerRequest)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid request: %s", err.Error()),
		})
		return
	}
	log.Debug(registerRequest)

	createdUser, er := services.UserClient.Register(registerRequest)
	if er != nil {
		c.JSON(er.Status(), gin.H{
			"message": fmt.Sprintf("Error creating user: %s", er.Error()),
		})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
