package clients

import (
	"cursos-ucc/model"
	error "cursos-ucc/utils/errors"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

type registerClient struct{}

type RegisterClientInterface interface {
	RegisterUserToCourse(userId int, courseId int) (model.Register, error.ApiError)
}

var (
	RegisterClient RegisterClientInterface
)

func init() {
	RegisterClient = &registerClient{}
}

func (r *registerClient) RegisterUserToCourse(userId int, courseId int) (model.Register, error.ApiError) {
	var register model.Register
	register.CourseID = courseId
	register.UserID = userId
	Db.Create(&register)
	return register, nil
}
