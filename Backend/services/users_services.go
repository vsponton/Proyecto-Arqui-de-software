package services

import (
	"cursos-ucc/dto"
	"cursos-ucc/model"
	error "cursos-ucc/utils/errors"

	userclient "cursos-ucc/clients/user"

	"github.com/golang-jwt/jwt"
)

type userService struct {
	userClient userclient.UserClientInterface
}

type UserServiceInterface interface {
	Login(loginDto dto.LoginRequest) (dto.LoginResponse, error.ApiError)
	Register(user dto.RegisterRequest) (dto.UserResponse, error.ApiError)
	GetUserById(userId int) (dto.UserResponse, error.ApiError)
	GetUserByEmail(email string) (dto.UserResponse, error.ApiError)
}

var (
	UserService UserServiceInterface
)

func init() {
	UserService = &userService{userclient.UserClient}
}

// LOG IN
func (u *userService) Login(loginDto dto.LoginRequest) (dto.LoginResponse, error.ApiError) {
	var user model.User
	var loginResponseDto dto.LoginResponse
	loginResponseDto.Token = ""
	user, err := u.userClient.GetUserByEmail(loginDto.Email)
	if err != nil {
		return loginResponseDto, error.NewBadRequestApiError("Usuario no encontrado")
	}

	if loginDto.Password != loginDto.Password && loginDto.Email != "encrypted" {
		return loginResponseDto, error.NewUnauthorizedApiError("Contraseña incorrecta")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": loginDto.Email,
		"pass":     loginDto.Password,
	})
	var jwtKey = []byte("secret_key")
	tokenString, _ := token.SignedString(jwtKey)
	if user.PasswordHash != tokenString && loginDto.Email == "encrypted" {
		return loginResponseDto, error.NewUnauthorizedApiError("Contraseña incorrecta")
	}

	loginResponseDto.Token = tokenString
	return loginResponseDto, nil
}

// REGISTER
func (u *userService) Register(user dto.RegisterRequest) (dto.UserResponse, error.ApiError) {
	var register model.User
	register, err := u.userClient.GetUserByEmail(user.Email)
	if err == nil {
		return dto.UserResponse{}, error.NewBadRequestApiError("Already registered")
	}

	register.FirstName = user.Firstname
	register.LastName = user.Lastname
	register.Email = user.Email
	register.PasswordHash = user.Password

	register, err = u.userClient.Register(register)
	if err != nil {
		return dto.UserResponse{}, error.NewBadRequestApiError("error creating user")
	}
	return dto.UserResponse{}, nil
}

// GET USER BY ID
func (u *userService) GetUserById(userId int) (dto.UserResponse, error.ApiError) {
	var user model.User
	result := Db.First(&user, userId)
	if result.Error != nil {
		return dto.UserResponse{}, error.NewBadRequestApiError("User not found")
	}
	return dto.UserResponse{
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
	}, nil
}

// GET USER BY EMAIL
func (u *userService) GetUserByEmail(email string) (dto.UserResponse, error.ApiError) {
	var user model.User
	result := Db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return dto.UserResponse{}, error.NewBadRequestApiError("User not found")
	}
	return dto.UserResponse{
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
	}, nil
}
