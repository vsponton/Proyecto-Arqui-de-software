package services

import (
	"cursos-ucc/dto"
	"cursos-ucc/model"
	error "cursos-ucc/utils/errors"

	//"cursos-ucc/db"
	courseclient "cursos-ucc/clients/course"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

type courseService struct {
	courseClient courseclient.CourseClientInterface
}

type CourseServiceInterface interface {
	GetCourses() (dto.CoursesResponse_Full, error.ApiError)
	GetCoursesByUserId(userId int) (dto.CoursesResponse_Full, error.ApiError)
	SearchCoursesByTitle(title string) (dto.CoursesResponse_Full, error.ApiError)
	SearchCoursesByCategory(category string) (dto.CoursesResponse_Full, error.ApiError)
	SearchCoursesByDescription(description string) (dto.CoursesResponse_Full, error.ApiError)
	CreateCourse(course dto.CourseResponse_Full) (dto.CourseResponse_Full, error.ApiError)
	UpdateCourse(course dto.CourseResponse_Full) (dto.CourseResponse_Full, error.ApiError)
	DeleteCourse(courseId int) error.ApiError
}

var (
	CourseService CourseServiceInterface
)

func init() {
	CourseService = &courseService{courseclient.CourseClient}
}

// DEVUELVE TODOS LOS CURSOS
// router.GET("/course", courseController.GetCourses)

func (c *courseService) GetCourses() (dto.CoursesResponse_Full, error.ApiError) {

	var coursesDto dto.CoursesResponse_Full

	courses, err := c.courseClient.GetCourses()

	if err != nil {
		return dto.CoursesResponse_Full{}, err
	}

	for _, course := range courses {
		var courseDto dto.CourseResponse_Full
		// como viene en el dto ---> como va en el model
		courseDto.ID_Course = course.ID
		courseDto.Title = course.Title
		courseDto.Description = course.Description
		courseDto.Category = course.Category
		courseDto.ImageURL = course.ImageURL
		courseDto.Duration = course.Duration
		courseDto.Requirements = course.Requirements

		coursesDto = append(coursesDto, courseDto)
	}

	return coursesDto, nil
}

func (c *courseService) GetCoursesByUserId(userId int) (dto.CoursesResponse_Full, error.ApiError) {

	var coursesDto dto.CoursesResponse_Full
	var courses model.Courses
	courses, err := c.courseClient.GetCoursesByUserId(userId)

	if err != nil {
		return nil, error.NewNotFoundApiError("Couldn't find courses by that user")
	}

	for _, course := range courses {
		var courseDto dto.CourseResponse_Full
		courseDto.ID_Course = course.ID
		courseDto.Title = course.Title
		courseDto.Description = course.Description
		courseDto.Category = course.Category
		courseDto.ImageURL = course.ImageURL
		courseDto.Duration = course.Duration
		courseDto.Requirements = course.Requirements

		coursesDto = append(coursesDto, courseDto)
	}

	return coursesDto, nil
}

func (c *courseService) SearchCoursesByTitle(title string) (dto.CoursesResponse_Full, error.ApiError) {

	var coursesDto dto.CoursesResponse_Full

	courses, _ := c.courseClient.SearchCoursesByTitle(title)

	for _, course := range courses {
		var courseDto dto.CourseResponse_Full
		courseDto.ID_Course = course.ID
		courseDto.Title = course.Title
		courseDto.Description = course.Description
		courseDto.Category = course.Category
		courseDto.ImageURL = course.ImageURL
		courseDto.Duration = course.Duration
		courseDto.Requirements = course.Requirements

		coursesDto = append(coursesDto, courseDto)
	}

	return coursesDto, nil
}

func (c *courseService) SearchCoursesByCategory(category string) (dto.CoursesResponse_Full, error.ApiError) {

	var coursesDto dto.CoursesResponse_Full

	courses, _ := c.courseClient.SearchCoursesByCategory(category)

	for _, course := range courses {
		var courseDto dto.CourseResponse_Full
		courseDto.ID_Course = course.ID
		courseDto.Title = course.Title
		courseDto.Description = course.Description
		courseDto.Category = course.Category
		courseDto.ImageURL = course.ImageURL
		courseDto.Duration = course.Duration
		courseDto.Requirements = course.Requirements

		coursesDto = append(coursesDto, courseDto)
	}

	return coursesDto, nil
}

func (c *courseService) SearchCoursesByDescription(description string) (dto.CoursesResponse_Full, error.ApiError) {

	var coursesDto dto.CoursesResponse_Full

	courses, _ := c.courseClient.SearchCoursesByDescription(description)

	for _, course := range courses {
		var courseDto dto.CourseResponse_Full
		courseDto.ID_Course = course.ID
		courseDto.Title = course.Title
		courseDto.Description = course.Description
		courseDto.Category = course.Category
		courseDto.ImageURL = course.ImageURL
		courseDto.Duration = course.Duration
		courseDto.Requirements = course.Requirements

		coursesDto = append(coursesDto, courseDto)
	}

	return coursesDto, nil
}

func (c *courseService) CreateCourse(course dto.CourseResponse_Full) (dto.CourseResponse_Full, error.ApiError) {
	result := Db.Create(&course)
	if result.Error != nil {
		return dto.CourseResponse_Full{}, error.NewInternalServerApiError("error creating course", result.Error)
	}
	return dto.CourseResponse_Full{}, nil

}

func (c *courseService) UpdateCourse(course dto.CourseResponse_Full) (dto.CourseResponse_Full, error.ApiError) {

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

func (c *courseService) DeleteCourse(courseId int) error.ApiError {
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
