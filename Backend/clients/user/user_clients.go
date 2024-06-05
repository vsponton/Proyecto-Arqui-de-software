package user

import (
	"cursos-ucc/model"
	"cursos-ucc/model"
	error "cursos-ucc/utils/errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

type userClient struct{}

type UserClientInterface interface {
	Login(ID_Course int64) (model.LoginResponse, error.ApiError)
	Register(user model.Register) (model.Users, error.ApiError)
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

func (u *userClient)Login(userID int64) (model.LoginResponse, error.ApiError) {
	var user model.LoginResponse

	user, err := u.userClient.GetCourseByIdUser(loginmodel.id_user)
	var loginResponsemodel model.LoginResponsemodel
	loginResponsemodel.UserId = -1
	if err != nil {
		return loginResponsemodel, e.NewBadRequestApiError("Usuario no encontrado")
	}
	if user.Password != loginmodel.Password && loginmodel.Username != "encrypted" {
		return loginResponsemodel, e.NewUnauthorizedApiError("Contraseña incorrecta")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": loginmodel.Username,
		"pass":     loginmodel.Password,
	})
	var jwtKey = []byte("secret_key")
	tokenString, _ := token.SignedString(jwtKey)
	if user.Password != tokenString && loginmodel.Username == "encrypted" {
		return loginResponsemodel, e.NewUnauthorizedApiError("Contraseña incorrecta")
	}

	loginResponsemodel.UserId = user.UserId
	loginResponsemodel.Token = tokenString
	log.Debug(loginResponsemodel)
	return loginResponsemodel, nil
}

func (u *userClient)Register(user model.Register) (model.Users, error.ApiError) {

	var register model.User
	//var reg model.RegisterRequest

	result := Db.Where("email = ? ", user.Email).First(&register)
	if result.Error == nil {
		return model.Register{}, error.NewBadRequestApiError("Already registered")
	}

	register.FirstName = user.Firstname
	register.LastName = user.Lastname
	register.Email = user.Email
	register.PasswordHash = user.Password

	
	result = Db.Create(&register)
	if result.Error != nil {
		return 
	}


	return registermodel, nil

}

// GET INFO ABOUT THE USER

func (u *userClient)GetUserById(id int64) (model.Users, error.ApiError) {
	var user model.Users
	result := Db.First(&user, id)
	if result.Error != nil {
		return nil, error.NewNotFoundApiError("???")
	}
	return user, nil
}

func (u *userClient)GetUserByEmail(email string) (model.Users, error.ApiError) {
	var user model.Users
	result := Db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, error.NewNotFoundApiError("???")
	}
	return user, nil
}



