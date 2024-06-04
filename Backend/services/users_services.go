package services

import (
	client "cursos-ucc/clients/user"
	"cursos-ucc/dto"
	"cursos-ucc/model"

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


////
