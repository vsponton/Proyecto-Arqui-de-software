package user

import (
	"cursos-ucc/model"
	error "cursos-ucc/utils/errors"

	log "cursos-ucc/logging"

	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

type userClient struct{}

type UserClientInterface interface {
	Login(ID_Course int64) (model.Users, error.ApiError)
	Register(user model.User) (model.Users, error.ApiError)
	GetUserById(userId int64) (model.Users, error.ApiError)
	GetUserByEmail(email string) (model.Users, error.ApiError)
}

var (
	UserClient UserClientInterface
)

func init() {
	UserClient = &userClient{}
}

// LOG IN

func (u *userClient) Login(userID int64) (model.Users, error.ApiError) {
	var user model.User

	

	user, err := u.userClient.GetCourseByIdUser(user.ID)
	var loginResponsemodel model.Users
	user.ID = -1
	if err != nil {
		return loginResponsemodel, e.NewBadRequestApiError("Usuario no encontrado")
	}
	if user.PasswordHash != user.PasswordHash && user.Email != "encrypted" {
		return loginResponsemodel, e.NewUnauthorizedApiError("Contraseña incorrecta")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Email,
		"pass":     user.PasswordHash,
	})
	var jwtKey = []byte("secret_key")
	tokenString, _ := token.SignedString(jwtKey)
	if user.PasswordHash != tokenString && user.Email == "encrypted" {
		return loginResponsemodel, e.NewUnauthorizedApiError("Wrong password")
	}

	loginResponsemodel.UserId = user.ID
	loginResponsemodel.Token = tokenString
	log.Debug(loginResponsemodel)
	return loginResponsemodel, nil
}

func (u *userClient) Register(user model.User) (model.Users, error.ApiError) {

	var register model.User
	//var reg model.RegisterRequest

	result := Db.Where("email = ? ", user.Email).First(&register)
	if result.Error == nil {
		return model.Users{}, error.NewBadRequestApiError("Already registered")
	}

	register.FirstName = user.FirstName
	register.LastName = user.LastName
	register.Email = user.Email
	register.PasswordHash = user.PasswordHash

	result = Db.Create(&register)
	if result.Error != nil {
		return model.Users{}, error.NewBadRequestApiError("error creating user")
	}

	return model.Users{}, nil

}

// GET INFO ABOUT THE USER

func (u *userClient) GetUserById(id int64) (model.Users, error.ApiError) {
	var user model.Users
	result := Db.First(&user, id)
	if result.Error != nil {
		return nil, error.NewNotFoundApiError("error getting user by id")
	}
	return user, nil
}

func (u *userClient) GetUserByEmail(email string) (model.Users, error.ApiError) {
	var user model.Users
	result := Db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, error.NewNotFoundApiError("error getting user by email")
	}
	return user, nil
}