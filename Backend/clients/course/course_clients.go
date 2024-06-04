package clients

import (
	"cursos-ucc/dto"
	//"cursos-ucc/logging/logging.go"

	error "cursos-ucc/utils/errors"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

type CourseClientInterface interface {
	GetCoursesByUserId(userId int) (dto.CoursesResponse_Full, error.ApiError)
	SearchCoursesByTitle(title string) (dto.CoursesResponse_Full, error.ApiError)
	SearchCoursesByCategory(category string) (dto.CoursesResponse_Full, error.ApiError)
	SearchCoursesByDescription(description string) (dto.CoursesResponse_Full, error.ApiError)
	CreateCourse(course dto.CourseResponse_Full) (dto.CourseResponse_Full, error.ApiError)
	UpdateCourse(course dto.CourseRequest_Registration) (dto.CourseResponse_Full, error.ApiError)
	DeleteCourse(courseId int) error.ApiError
}

type courseClient struct{}

var CourseClient CourseClientInterface = &courseClient{}

// func (c *courseClient) GetCoursesByUserId(userId int) (dto.CoursesResponse_Full, error.ApiError)

func (c *courseClient) SearchCoursesByTitle(title string) (dto.CoursesResponse_Full, error.ApiError) {

	var course dto.CoursesResponse_Full

	result := Db.Where("title LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").Find(&course)

	if result.Error != nil {
		return nil, result.Error
	}
	return course, nil

	log.Debug("Course", course)

	return dto.CoursesResponse_Full{}, nil
}

func (c *courseClient) SearchCoursesByCategory(category string) (dto.CoursesResponse_Full, error.ApiError) {

	var course dto.CoursesResponse_Full

	result := Db.Where("title LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").Find(&course)

	if result.Error != nil {
		return nil, result.Error
	}
	return course, nil

	log.Debug("Course", course)

	return dto.CoursesResponse_Full{}, nil
}

func (c *courseClient) SearchCoursesByDescription(description string) (dto.CoursesResponse_Full, error.ApiError) {

	var course dto.CoursesResponse_Full

	result := Db.Where("title LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").Find(&course)

	if result.Error != nil {
		return nil, result.Error
	}
	return course, nil

	log.Debug("Course", course)

	return dto.CoursesResponse_Full{}, nil
}

func (c *courseClient) CreateCourse(course dto.CourseRequest_Registration) (dto.CoursesResponse_Full, error.ApiError) {
	result := Db.Create(&course)
	if result.Error != nil {
		//return dto.CourseResponse_Full{}, errors.NewInternalServerError("error creating course")
	}
	return dto.CoursesResponse_Full{}, nil

}

func (c *courseClient) UpdateCourse(course dto.CourseRequest_Registration) (dto.CourseResponse_Full, error.ApiError) {

	var existingCourse dto.CourseRequest_Registration

	// Verificar si el curso existe
	if err := Db.Where("id = ?", ID_Course).First(&existingCourse).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return dto.CourseResponse_Full{}, errors.NewNotFoundError("course not found")
		}
		return dto.CourseResponse_Full{}, errors.NewInternalServerError("error finding course")
	}

	// Actualizar el curso
	result := Db.Model(&existingCourse).Updates(course)
	if result.Error != nil {
		return dto.CourseResponse_Full{}, errors.NewInternalServerError("error updating course")
	}

	// Retornar el curso actualizado
	return existingCourse.ToResponse(), nil
}

func (c *courseClient) DeleteCourse(courseId int) error.ApiError {
	var course dto.CourseRequest_Registration

	// Verificar si el curso existe
	if err := Db.Where("id = ?", courseId).First(&course).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.NewNotFoundError("course not found")
		}
		return errors.NewInternalServerError("error finding course")
	}

	// Eliminar el curso
	if err := Db.Delete(&course).Error; err != nil {
		return errors.NewInternalServerError("error deleting course")
	}
	return nil
}

/*
func InsertSubscription(userID int64, courseID int64) error {
	var subscription model.Subscription
	result := db.Where("user_id = ? AND course_id = ?", userID, courseID).First(&subscription)
	if result.Error == nil {
		return fmt.Errorf("user %d is already subscribed to course %d", userID, courseID)
	}

	subscription = model.Subscription{
		UserID:       userID,
		CourseID:     courseID,
		CreationDate: time.Now().UTC(),
		LastUpdated:  time.Now().UTC(),
	}

	result = db.Create(&subscription)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
*/
