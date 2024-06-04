package services

import (
	clients "cursos-ucc/clients/course"

	"cursos-ucc/dto"

	"cursos-ucc/logging"

	error "cursos-ucc/utils/errors"
)

type courseService struct{}

type CourseServiceInterface interface {
	GetCourseByIdUser(userId int) (dto.CoursesResponse_Full, error.ApiError)
	SearchByTitle(title string) (dto.CoursesResponse_Full, error.ApiError)
	SearchByCategory(category string) (dto.CoursesResponse_Full, error.ApiError)
	SearchByDescription(description string) (dto.CoursesResponse_Full, error.ApiError)
	PostCourse(course dto.CourseResponse_Full) (dto.CourseResponse_Full, error.ApiError)
	PutCourse(course dto.CourseRequest_Registration) (dto.CourseResponse_Full, error.ApiError)
	DeleteCourse(courseId int) error.ApiError
}

var CourseService CourseServiceInterface = &courseService{}

func (s *courseService) GetCourseByIdUser(userId int) (dto.CoursesResponse_Full, error.ApiError) {
	courses, err := clients.CoursesClient.GetCoursesByUserId(userId) // clients.CourseService ????????
	if err != nil {
		logging.Log.Error("Error fetching courses for user ID: ", userId, " Error: ", err)
		return nil, error.NewInternalServerApiError("Error fetching courses", err)
	}
	return courses, nil
}

func (s *courseService) SearchByTitle(title string) (dto.CoursesResponse_Full, error.ApiError) {
	courses, err := clients.CoursesClient.SearchCoursesByTitle(title)
	if err != nil {
		logging.Log.Error("Error searching courses by title: ", title, " Error: ", err)
		return nil, error.NewNotFoundApiError("No courses found with title " + title)
	}
	return courses, nil
}

func (s *courseService) SearchByCategory(category string) (dto.CoursesResponse_Full, error.ApiError) {
	courses, err := clients.CoursesClient.SearchCoursesByCategory(category)
	if err != nil {
		logging.Log.Error("Error searching courses by category: ", category, " Error: ", err)
		return nil, error.NewNotFoundApiError("No courses found with category " + category)
	}
	return courses, nil
}

func (s *courseService) SearchByDescription(description string) (dto.CoursesResponse_Full, error.ApiError) {
	courses, err := clients.CoursesClient.SearchCoursesByDescription(description)
	if err != nil {
		logging.Log.Error("Error searching courses by description: ", description, " Error: ", err)
		return nil, error.NewNotFoundApiError("No courses found with description " + description)
	}
	return courses, nil
}

func (s *courseService) PostCourse(course dto.CourseResponse_Full) (dto.CourseResponse_Full, error.ApiError) {
	newCourse, err := clients.CoursesClient.CreateCourse(course)
	if err != nil {
		logging.Log.Error("Error creating course: ", course, " Error: ", err)
		return dto.CourseResponse_Full{}, error.NewInternalServerApiError("Error creating course", err)
	}
	return newCourse, nil
}

func (s *courseService) PutCourse(course dto.CourseRequest_Registration) (dto.CourseResponse_Full, error.ApiError) {
	updatedCourse, err := clients.CoursesClient.UpdateCourse(course)
	if err != nil {
		logging.Log.Error("Error updating course: ", course, " Error: ", err)
		return dto.CourseResponse_Full{}, error.NewInternalServerApiError("Error updating course", err)
	}
	return updatedCourse, nil
}

func (s *courseService) DeleteCourse(courseId int) error.ApiError {
	err := clients.CoursesClient.DeleteCourse(courseId)
	if err != nil {
		logging.Log.Error("Error deleting course with ID: ", courseId, " Error: ", err)
		return error.NewInternalServerApiError("Error deleting course", err)
	}
	return nil
}
