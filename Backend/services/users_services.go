package services

import (
	"cursos-ucc/dto"
	"cursos-ucc/model"
	error "cursos-ucc/utils/errors"

	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

type userClient struct{}

type UserClientInterface interface {
	Login(loginDto dto.LoginRequest) (dto.LoginResponse, error.ApiError)
	Register(user dto.RegisterRequest) (dto.UserResponse, error.ApiError)
	GetUserById(userId int64) (dto.UserResponse, error.ApiError)
	GetUserByEmail(email string) (dto.UserResponse, error.ApiError)
}

var (
	UserClient UserClientInterface
)

func init() {
	UserClient = &userClient{}
}

// LOG IN

func (u *userClient) Login(loginDto dto.LoginRequest) (dto.LoginResponse, error.ApiError) {

	var loginResponseDto dto.LoginResponse
	loginResponseDto.Token = ""
	if err != nil {
		return loginResponseDto, error.NewBadRequestApiError("Usuario no encontrado")
	}

	// llamada al back para pedirle el usuario por usuario, que nos de la contraseña y comparar

	if loginDto.Password != loginDto.Password && loginDto.Email != "encrypted" {
		return loginResponseDto, error.NewUnauthorizedApiError("Contraseña incorrecta")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": loginDto.Email,
		"pass":     loginDto.Password,
	})
	var jwtKey = []byte("secret_key")
	tokenString, _ := token.SignedString(jwtKey)
	if User.Password != tokenString && loginDto.Email == "encrypted" {
		return loginResponseDto, error.NewUnauthorizedApiError("Contraseña incorrecta")
	}

	//	loginResponseDto.UserId = user.UserId
	loginResponseDto.Token = tokenString
	//	log.Debug(loginResponseDto)
	return loginResponseDto, nil
}

func (u *userClient) Register(user dto.RegisterRequest) (dto.UserResponse, error.ApiError) {

	var register model.User
	//var reg dto.RegisterRequest

	result := Db.Where("email = ? ", user.Email).First(&register)
	if result.Error == nil {
		return dto.RegisterResponse{}, error.NewBadRequestApiError("Already registered")
	}

	register.FirstName = user.Firstname
	register.LastName = user.Lastname
	register.Email = user.Email
	register.PasswordHash = user.Password

	result = Db.Create(&register)
	if result.Error != nil {
		return
	}

	return Register, nil

}

// GET INFO ABOUT THE USER

func (u *userClient) GetUserById(id int64) (dto.UserResponse, error.ApiError) {
	var user dto.UserResponse
	result := Db.First(&user, id)
	if result.Error != nil {
		return nil, error.NewNotFoundApiError("???")
	}
	return user, nil
}

func (u *userClient) GetUserByEmail(email string) (dto.UserResponse, error.ApiError) {
	var user dto.UserResponse
	result := Db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, error.NewNotFoundApiError("???")
	}
	return user, nil
}

/*
package services

import (
	client "cursos-ucc/clients/user"
	"cursos-ucc/dto"

	"github.com/golang-jwt/jwt"

	error "cursos-ucc/utils/errors"

	log "github.com/sirupsen/logrus"
)

type userService struct {
	userClient client.UserClientInterface
}

type userServiceInterface interface {
	GetUserById(id int) (dto.UserResponse, error.ApiError)
	GetUserByEmail() (dto.UserResponse, error.ApiError)
	//Login(loginDto dto.LoginRequest) (dto.LoginResponse, error.ApiError)
}

var (
	UserService userServiceInterface
)

func initUserService(userClient client.UserServiceInterface) userServiceInterface {
	service := new(userService)
	service.userClient = userClient
	return service
}

func init() {
	UserService = initUserService(client.UserClient)
}
*/
