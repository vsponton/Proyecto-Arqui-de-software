package user

import (
	"cursos-ucc/model"
	"fmt"
	"github.com/jinzhu/gorm"
	
)

var Db *gorm.DB

type userClient struct{}

type UserClientInterface interface {
	Login(ID_Course int64) model.Register
	Register(ID_User int64) model.Register
}

var (
	UserClient UserClientInterface
)

func init() {
	UserClient = &userClient{}
}

func Login(userID int64, courseID int64) error {


}

func Register(userID int64, courseID int64) error {

	var login model.Register
	result := Db.Where("user_id = ? AND course_id = ?", userID, courseID).First(&register)
	if result.Error == nil {
		return fmt.Errorf("user %d is already subscribed to course %d", userID, courseID)
	}

	register = model.Register{
		UserID:       userID,
		CourseID:     courseID,
		
	}

	result = Db.Create(&register)
	if result.Error != nil {
		return result.Error
	}

	return nil

}
