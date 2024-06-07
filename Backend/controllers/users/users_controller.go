package usersController

import (
	"cursos-ucc/dto"
	service "cursos-ucc/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {

	// log.Debug("User id: " + c.Param("id"))

	//var UsersResponse dto.UsersResponse

	id, _ := strconv.Atoi(c.Param("user_id"))

	var userDto dto.UserResponse

	//UsersResponse, err := service.UserService.GetUserById(id)
	userDto, err := service.UserService.GetUserById(id)

	if err != nil {
		// log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, userDto)
}

func GetUserByEmail(c *gin.Context) {
	var email string
	email = c.Param(email)

	var UsersResponse dto.UserResponse

	UsersResponse, err := service.UserService.GetUserByEmail(email)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, UsersResponse)
}

// GenerateToken handles the token generation.
func Login(c *gin.Context) {

	var request dto.LoginRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := service.UserService.Login(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, user)
}

///////////////////

/*

package users

import (
    "net/http"
    usersDomain "backend/domain/users"
    usersService "backend/services/users"
    "github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
    var loginRequest usersDomain.LoginRequest
    if err := context.BindJSON(&loginRequest); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }
    response, err := usersService.Login(loginRequest)
    if err != nil {
        context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }
    context.JSON(http.StatusOK, response)
}

// CreateUser handles user creation requests.
func CreateUser(context *gin.Context) {
    var user usersDomain.User
    if err := context.BindJSON(&user); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    if err := usersService.CreateUser(user); err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// GenerateToken handles the token generation.
func GenerateToken(c *gin.Context) {
    var request usersDomain.TokenRequest
    if err := c.BindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    tokenString, err := usersService.GenerateJWT(request.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
*/
