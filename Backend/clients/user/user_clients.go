package user

import (
	"cursos-ucc/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func GetUserById(id int64) (model.User, error) {
	var user model.User
	result := Db.First(&user, id)
	if result.Error != nil {
		return model.User{}, fmt.Errorf("not found user with ID: %d", id)
	}
	return user, nil
}

func GetUserByEmail(email string) (model.User, error) {
	var user model.User
	result := Db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return model.User{}, fmt.Errorf("not found user with email: %s", email)
	}
	return user, nil
}


