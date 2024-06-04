package user

import (
	"cursos-ucc/model"
	"fmt"
	"time"
	"github.com/jinzhu/gorm"
	
)

var Db *gorm.DB

type userClient struct{}

type UserClientInterface interface {
	Login(ID_Course int64) model.Subscription
	Register(ID_User int64) model.Subscription
}

var (
	UserClient UserClientInterface
)

func init() {
	UserClient = &userClient{}
}

func Login(userID int64, courseID int64) error {

	//////////////

}

func Register(userID int64, courseID int64) error {

	var login model.Subscription
	result := Db.Where("user_id = ? AND course_id = ?", userID, courseID).First(&subscription)
	if result.Error == nil {
		return fmt.Errorf("user %d is already subscribed to course %d", userID, courseID)
	}

	subscription = model.Subscription{
		UserID:       userID,
		CourseID:     courseID,
		CreationDate: time.Now().UTC(),
		LastUpdated:  time.Now().UTC(),
	}

	result = Db.Create(&subscription)
	if result.Error != nil {
		return result.Error
	}

	return nil

}
