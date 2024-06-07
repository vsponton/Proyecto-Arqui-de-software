// package services

// import (
// 	"cursos-ucc/dto"
// 	"cursos-ucc/model"
// 	error "cursos-ucc/utils/errors"

// 	"github.com/golang-jwt/jwt"
// 	"github.com/jinzhu/gorm"
// )

// var Db *gorm.DB

// type userClient struct {
// 	Db *gorm.DB
// }

// //type userClient struct{}
// type registerClient struct{}

// type UserClientInterface interface {
// 	Login(loginDto dto.LoginRequest) (dto.LoginResponse, error.ApiError)
// 	Register(user dto.RegisterRequest) (dto.UserResponse, error.ApiError)
// 	GetUserById(userId int64) (dto.UserResponse, error.ApiError)
// 	GetUserByEmail(email string) (dto.UserResponse, error.ApiError)
// }

// var (
// 	UserClient UserClientInterface
// )

// func init() {
// 	UserClient = &userClient{}
// }

// // LOG IN

// func (u *userClient) Login(loginDto dto.LoginRequest) (dto.LoginResponse, error.ApiError) {

// 	var user model.User

// 	var loginResponseDto dto.LoginResponse
// 	loginResponseDto.Token = ""
// 	err := Db.Where("email = ?", loginDto.Email).First(&user).Error
// 	if err != nil {
// 		return loginResponseDto, error.NewBadRequestApiError("Usuario no encontrado")
// 	}

// 	// llamada al back para pedirle el usuario por usuario, que nos de la contraseña y comparar

// 	if loginDto.Password != loginDto.Password && loginDto.Email != "encrypted" {
// 		return loginResponseDto, error.NewUnauthorizedApiError("Contraseña incorrecta")
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"username": loginDto.Email,
// 		"pass":     loginDto.Password,
// 	})
// 	var jwtKey = []byte("secret_key")
// 	tokenString, _ := token.SignedString(jwtKey)
// 	if user.PasswordHash != tokenString && loginDto.Email == "encrypted" {
// 		return loginResponseDto, error.NewUnauthorizedApiError("Contraseña incorrecta")
// 	}

// 	//	loginResponseDto.UserId = user.UserId
// 	loginResponseDto.Token = tokenString
// 	//	log.Debug(loginResponseDto)
// 	return loginResponseDto, nil
// }

// func (u *userClient) Register(user dto.RegisterRequest) (dto.UserResponse, error.ApiError) {

// 	var register model.User
// 	//var hello dto.RegisterRequest

// 	result := Db.Where("email = ? ", user.Email).First(&register)
// 	if result.Error == nil {
// 		return dto.UserResponse{}, error.NewBadRequestApiError("Already registered")
// 	}

// 	register.FirstName = user.Firstname
// 	register.LastName = user.Lastname
// 	register.Email = user.Email
// 	register.PasswordHash = user.Password

// 	result = Db.Create(&register)
// 	if result.Error != nil {
// 		return dto.UserResponse{}, error.NewBadRequestApiError("error creating user")
// 	}

// 	return dto.UserResponse{}, nil

// }

// // GET INFO ABOUT THE USER

// func (u *userClient) GetUserById(id int64) (dto.UserResponse, error.ApiError) {
// 	var user dto.UserResponse
// 	result := Db.First(&user, id)
// 	if result.Error != nil {
// 		return dto.UserResponse{}, error.NewBadRequestApiError("User not found")
// 	}
// 	return user, nil
// }

// func (u *userClient) GetUserByEmail(email string) (dto.UserResponse, error.ApiError) {

// 	var user dto.UserResponse

// 	result := Db.Where("email = ?", "%"+email+"%").First(&user)

// 	if result.Error != nil {

// 		return dto.UserResponse{}, error.NewBadRequestApiError("User not found")
// 	}
// 	return user, nil
// }

package services

import (
	"cursos-ucc/dto"
	"cursos-ucc/model"
	error "cursos-ucc/utils/errors"

	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

type userClient struct {
	Db *gorm.DB
}

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
	UserClient = &userClient{Db: Db}
}

// LOG IN
func (u *userClient) Login(loginDto dto.LoginRequest) (dto.LoginResponse, error.ApiError) {
	var user model.User
	var loginResponseDto dto.LoginResponse
	loginResponseDto.Token = ""
	err := u.Db.Where("email = ?", loginDto.Email).First(&user).Error
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
func (u *userClient) Register(user dto.RegisterRequest) (dto.UserResponse, error.ApiError) {
	var register model.User
	result := u.Db.Where("email = ?", user.Email).First(&register)
	if result.Error == nil {
		return dto.UserResponse{}, error.NewBadRequestApiError("Already registered")
	}

	register.FirstName = user.Firstname
	register.LastName = user.Lastname
	register.Email = user.Email
	register.PasswordHash = user.Password

	result = u.Db.Create(&register)
	if result.Error != nil {
		return dto.UserResponse{}, error.NewBadRequestApiError("error creating user")
	}

	return dto.UserResponse{}, nil
}

// GET USER BY ID
func (u *userClient) GetUserById(userId int64) (dto.UserResponse, error.ApiError) {
	var user model.User
	result := u.Db.First(&user, userId)
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
func (u *userClient) GetUserByEmail(email string) (dto.UserResponse, error.ApiError) {
	var user model.User
	result := u.Db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return dto.UserResponse{}, error.NewBadRequestApiError("User not found")
	}
	return dto.UserResponse{
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
	}, nil
}
