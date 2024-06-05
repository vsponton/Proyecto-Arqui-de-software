package services

import (
	"cursos-ucc/dto"
	error "cursos-ucc/utils/errors"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

type courseClient struct{}

type CourseClientInterface interface {
	GetCoursesByUserId(userId int) (dto.CoursesResponse_Full, error.ApiError)
	SearchCoursesByTitle(title string) (dto.CoursesResponse_Full, error.ApiError)
	SearchCoursesByCategory(category string) (dto.CoursesResponse_Full, error.ApiError)
	SearchCoursesByDescription(description string) (dto.CoursesResponse_Full, error.ApiError)
	CreateCourse(course dto.CourseResponse_Full) (dto.CourseResponse_Full, error.ApiError)
	UpdateCourse(course dto.CourseResponse_Full) (dto.CourseResponse_Full, error.ApiError)
	DeleteCourse(courseId int) error.ApiError
}

var (
	CourseClient CourseClientInterface
)

func init() {
	CourseClient = &courseClient{}
}

func (c *courseClient) GetCoursesByUserId(userId int) (dto.CoursesResponse_Full, error.ApiError) {
	var course dto.CoursesResponse_Full

	result := Db.Where("id=?", userId).Find(&course)

	if result.Error != nil {
		return nil, error.NewNotFoundApiError("???")
	}
	return course, nil
}

func (c *courseClient) SearchCoursesByTitle(title string) (dto.CoursesResponse_Full, error.ApiError) {

	var course dto.CoursesResponse_Full

	result := Db.Where("title LIKE ?", "%"+title+"%").Find(&course)

	if result.Error != nil {
		return nil, error.NewNotFoundApiError("Error!")
	}
	return course, nil
}

func (c *courseClient) SearchCoursesByCategory(category string) (dto.CoursesResponse_Full, error.ApiError) {

	var course dto.CoursesResponse_Full

	result := Db.Where("category LIKE ?", "%"+category+"%").Find(&course)

	if result.Error != nil {
		return nil, error.NewNotFoundApiError("???")
	}
	return course, nil
}

func (c *courseClient) SearchCoursesByDescription(description string) (dto.CoursesResponse_Full, error.ApiError) {

	var course dto.CoursesResponse_Full

	result := Db.Where("description LIKE ?", "%"+description+"%").Find(&course)

	if result.Error != nil {
		return nil, error.NewNotFoundApiError("???")
	}
	return course, nil
}

func (c *courseClient) CreateCourse(course dto.CourseResponse_Full) (dto.CourseResponse_Full, error.ApiError) {
	result := Db.Create(&course)
	if result.Error != nil {
		return dto.CourseResponse_Full{}, error.NewInternalServerApiError("error creating course", result.Error)
	}
	return dto.CourseResponse_Full{}, nil

}

func (c *courseClient) UpdateCourse(course dto.CourseResponse_Full) (dto.CourseResponse_Full, error.ApiError) {

	var existingCourse dto.CourseResponse_Full

	// Verificar si el curso existe
	if err := Db.Where("id = ?", course.ID_Course).First(&existingCourse).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return dto.CourseResponse_Full{}, error.NewNotFoundApiError("course not found")
		}
		return dto.CourseResponse_Full{}, error.NewInternalServerApiError("error finding course", err)
	}

	// Actualizar el curso
	result := Db.Model(&existingCourse).Updates(course)
	if result.Error != nil {
		return dto.CourseResponse_Full{}, error.NewInternalServerApiError("error updating course", result.Error)
	}

	// Retornar el curso actualizado
	return dto.CourseResponse_Full{}, nil
}

func (c *courseClient) DeleteCourse(courseId int) error.ApiError {
	var course dto.CourseRequest_Registration

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

/*
package services

import (
	clients "cursos-ucc/clients/course"

	"cursos-ucc/dto"

	"cursos-ucc/logging"

	error "cursos-ucc/utils/errors"
)

type courseService struct{
	courseClient clients.CourseClientInterface
}

type CourseServiceInterface interface {
	GetCourseByIdUser(userId int) (dto.CoursesResponse_Full, error.ApiError)
	SearchByTitle(title string) (dto.CoursesResponse_Full, error.ApiError)
	SearchByCategory(category string) (dto.CoursesResponse_Full, error.ApiError)
	SearchByDescription(description string) (dto.CoursesResponse_Full, error.ApiError)
	PostCourse(course dto.CourseResponse_Full) (dto.CourseResponse_Full, error.ApiError)
	PutCourse(course dto.CourseResponse_Full) (dto.CourseResponse_Full, error.ApiError)
	DeleteCourse(courseId int) error.ApiError
}

var (
	CourseService CourseServiceInterface
)

func initCourseService(courseClient clients.CourseClientInterface) CourseServiceInterface {
	service := new(courseService)
	service.courseClient = courseClient
	return service
}

func init() {
	CourseService = initCourseService(clients.CourseClient)
}

func (s *courseService) GetCourseByIdUser(userId int) (dto.CoursesResponse_Full, error.ApiError) {
	courses, err := s.courseClient.GetCoursesByUserId(userId) // clients.CourseService ????????
	if err != nil {
		logging.Log.Error("Error fetching courses for user ID: ", userId, " Error: ", err)
		return nil, error.NewInternalServerApiError("Error fetching courses", err)
	}
	return courses, nil
}

func (s *courseService) SearchByTitle(title string) (dto.CoursesResponse_Full, error.ApiError) {
	courses, err := s.courseClient.SearchCoursesByTitle(title)
	if err != nil {
		logging.Log.Error("Error searching courses by title: ", title, " Error: ", err)
		return nil, error.NewNotFoundApiError("No courses found with title " + title)
	}
	return courses, nil
}

func (s *courseService) SearchByCategory(category string) (dto.CoursesResponse_Full, error.ApiError) {
	courses, err := s.courseClient.SearchCoursesByCategory(category)
	if err != nil {
		logging.Log.Error("Error searching courses by category: ", category, " Error: ", err)
		return nil, error.NewNotFoundApiError("No courses found with category " + category)
	}
	return courses, nil
}

func (s *courseService) SearchByDescription(description string) (dto.CoursesResponse_Full, error.ApiError) {
	courses, err := s.courseClient.SearchCoursesByDescription(description)
	if err != nil {
		logging.Log.Error("Error searching courses by description: ", description, " Error: ", err)
		return nil, error.NewNotFoundApiError("No courses found with description " + description)
	}
	return courses, nil
}

func (s *courseService) PostCourse(course dto.CourseResponse_Full) (dto.CourseResponse_Full, error.ApiError) {
	newCourse, err := s.courseClient.CreateCourse(course)
	if err != nil {
		logging.Log.Error("Error creating course: ", course, " Error: ", err)
		return dto.CourseResponse_Full{}, error.NewInternalServerApiError("Error creating course", err)
	}
	return newCourse, nil
}

func (s *courseService) PutCourse(course dto.CourseResponse_Full) (dto.CourseResponse_Full, error.ApiError) {
	updatedCourse, err := s.courseClient.UpdateCourse(course)
	if err != nil {
		logging.Log.Error("Error updating course: ", course, " Error: ", err)
		return dto.CourseResponse_Full{}, error.NewInternalServerApiError("Error updating course", err)
	}
	return updatedCourse, nil
}

func (s *courseService) DeleteCourse(courseId int) error.ApiError {
	err := s.courseClient.DeleteCourse(courseId)
	if err != nil {
		logging.Log.Error("Error deleting course with ID: ", courseId, " Error: ", err)
		return error.NewInternalServerApiError("Error deleting course", err)
	}
	return nil
}
*/
