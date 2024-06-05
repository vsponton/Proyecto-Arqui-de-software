package clients

import (
	"cursos-ucc/model"
	error "cursos-ucc/utils/errors"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

type courseClient struct{}

type CourseClientInterface interface {
	GetCoursesByUserId(userId int) (model.Courses, error.ApiError)
	SearchCoursesByTitle(title string) (model.Courses, error.ApiError)
	SearchCoursesByCategory(category string) (model.Courses, error.ApiError)
	SearchCoursesByDescription(description string) (model.Courses, error.ApiError)
	CreateCourse(course model.Course) (model.Course, error.ApiError)
	UpdateCourse(course model.Course) (model.Course, error.ApiError)
	DeleteCourse(courseId int) error.ApiError
}


var (
	CourseClient CourseClientInterface
)

func init(){
	CourseClient = &courseClient{}
}

func (c *courseClient) GetCoursesByUserId(userId int) (model.Courses, error.ApiError){
	var course model.Courses

	result := Db.Where("id=?", userId).Find(&course)

	if result.Error != nil {
		return nil, error.NewNotFoundApiError("???")
	}
	return course, nil
}

func (c *courseClient) SearchCoursesByTitle(title string) (model.Courses, error.ApiError) {

	var course model.Courses

	result := Db.Where("title LIKE ?", "%"+title+"%").Find(&course)

	if result.Error != nil {
		return nil, error.NewNotFoundApiError("Error!")
	}
	return course, nil
}

func (c *courseClient) SearchCoursesByCategory(category string) (model.Courses, error.ApiError) {

	var course model.Courses

	result := Db.Where("category LIKE ?", "%"+category+"%").Find(&course)

	if result.Error != nil {
		return nil, error.NewNotFoundApiError("???")
	}
	return course, nil
}

func (c *courseClient) SearchCoursesByDescription(description string) (model.Courses, error.ApiError) {

	var course model.Courses

	result := Db.Where("description LIKE ?", "%"+description+"%").Find(&course)

	if result.Error != nil {
		return nil, error.NewNotFoundApiError("???")
	}
	return course, nil
}

func (c *courseClient) CreateCourse(course model.Course) (model.Course, error.ApiError) {
	result := Db.Create(&course)
	if result.Error != nil {
		return model.Course{}, error.NewInternalServerApiError("error creating course", result.Error)
	}
	return model.Course{}, nil

}

func (c *courseClient) UpdateCourse(course model.Course) (model.Course, error.ApiError) {

	var existingCourse model.Course

	// Verificar si el curso existe
	if err := Db.Where("id = ?", course.ID).First(&existingCourse).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return model.Course{}, error.NewNotFoundApiError("course not found")
		}
		return model.Course{}, error.NewInternalServerApiError("error finding course", err)
	}

	// Actualizar el curso
	result := Db.Model(&existingCourse).Updates(course)
	if result.Error != nil {
		return model.Course{}, error.NewInternalServerApiError("error updating course", result.Error)
	}

	// Retornar el curso actualizado
	return model.Course{}, nil
}

func (c *courseClient) DeleteCourse(courseId int) error.ApiError {
	var course model.Course

	// Verificar si el curso existe
	if err := Db.Where("id = ?", courseId).First(&course).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return error.NewNotFoundApiError("course not found")
		}
		return error.NewInternalServerApiError("error finding course", err)
	}

	// Eliminar el curso
	if err := Db.Delete(&course).Error; err != nil {
		return error.NewInternalServerApiError("error deleting course", err)
	}
	return nil
}