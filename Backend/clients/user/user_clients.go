package user

import (
	"cursos-ucc/model"
	error "cursos-ucc/utils/errors"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

type userClient struct{}

type UserClientInterface interface {
	Register(user model.User) (model.User, error.ApiError)
	GetUserById(userId int) (model.User, error.ApiError)
	GetUserByEmail(email string) (model.User, error.ApiError)
}

var (
	UserClient UserClientInterface
)

func init() {
	UserClient = &userClient{}
}

func (u *userClient) Register(user model.User) (model.User, error.ApiError) {

	var register model.User
	//var reg model.RegisterRequest

	result := Db.Where("email = ? ", user.Email).First(&register)
	if result.Error == nil {
		return model.User{}, error.NewBadRequestApiError("Already registered")
	}

	register.FirstName = user.FirstName
	register.LastName = user.LastName
	register.Email = user.Email
	register.PasswordHash = user.PasswordHash

	result = Db.Create(&register)
	if result.Error != nil {
		return model.User{}, error.NewBadRequestApiError("error creating user")
	}

	return model.User{}, nil

}

// GET INFO ABOUT THE USER

func (u *userClient) GetUserById(id int) (model.User, error.ApiError) {
	var user model.User
	result := Db.First(&user, id)
	if result.Error != nil {
		return model.User{}, error.NewNotFoundApiError("error getting user by id")
	}
	return user, nil
}

func (u *userClient) GetUserByEmail(email string) (model.User, error.ApiError) {
	var user model.User
	result := Db.Where("email = ?", email).Find(&user)
	if result.Error != nil {
		return model.User{}, error.NewNotFoundApiError("error getting user by email")
	}
	return user, nil
}
