package services

import (
	"cursos-ucc/model"

	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"

	error "cursos-ucc/utils/errors"
)

type LoginService struct{}

type loginServiceInterface interface {
	Login() (dto.LoginRequest, error.ApiError)
}

func (s *userService) Login(loginDto dto.LoginRequest) (dto.LoginResponse, error.ApiError) {

	var user model.Subscription
	user, err := s.userClient.GetUserByUsername(loginDto.Username)
	var loginResponseDto dto.LoginResponse
	loginResponseDto.UserId = -1
	if err != nil {
		return loginResponseDto, error.NewBadRequestApiError("Usuario no encontrado")
	}
	if user.Password != loginDto.Password && loginDto.Email != "encrypted" {
		return loginResponseDto, error.NewUnauthorizedApiError("Contraseña incorrecta")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": loginDto.Email,
		"pass":     loginDto.Password,
	})
	var jwtKey = []byte("secret_key")
	tokenString, _ := token.SignedString(jwtKey)
	if user.Password != tokenString && loginDto.Email == "encrypted" {
		return loginResponseDto, error.NewUnauthorizedApiError("Contraseña incorrecta")
	}

	loginResponseDto.UserId = user.UserId
	loginResponseDto.Token = tokenString
	log.Debug(loginResponseDto)
	return loginResponseDto, nil
}

type registerService struct{}

type RegisterServiceInterface interface {
	InsertUser(dto.RegisterRequest) (dto.RegisterResponse, error.ApiError)
}

var RegisterService RegisterServiceInterface = &registerService{}

func (s *registerService) Register(registerRequest dto.RegisterRequest) (dto.RegisterResponse, error.ApiError) {
	var user model.Users
	// Simulación de inserción del usuario en la base de datos.
	user.FirstName = registerRequest.Firstname
	user.LastName = registerRequest.Lastname
	user.Email = registerRequest.Email
	user.Password = registerRequest.Password

	// Aquí deberías insertar el usuario en la base de datos y obtener el ID generado
	userId := 1 // Simulación de un ID de usuario generado

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"firstname": registerRequest.FirstName,
		"lastname":  registerRequest.LastName,
		"email":     registerRequest.Email,
	})
	var jwtKey = []byte("secret_key")
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return dto.RegisterResponse{}, error.NewInternalServerApiError("Error generating token", err)
	}

	createdUser := dto.RegisterResponse{
		Token: tokenString,
	}

	log.Debug(createdUser)
	return createdUser, nil
}
