package clients

import (
	"cursos-ucc/dto"
	"cursos-ucc/error"
	"cursos-ucc/model"
	"fmt"
)

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

func ConnectDatabase() error {
	dsn := "root:root@tcp(localhost:3306)/coursesplatform?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	err = db.AutoMigrate(&model.User{}, &model.Course{}, &model.Subscription{})
	if err != nil {
		return fmt.Errorf("failed to auto-migrate: %w", err)
	}

	return nil
}

// func (c *courseClient) GetCoursesByUserId(userId int) (dto.CoursesResponse_Full, error.ApiError)

func GetCoursesByUserId(query string) ([]model.Course, error) {
	var courses []model.Course
	result := db.Where("title LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}

func SelectCourseByIDUser(id int64) (model.Course, error) {
	var course dto.CoursesResponse_Full
	result := db.First(&course, id)
	if result.Error != nil {
		return model.Course{}, fmt.Errorf("not found course with ID: %d", id)
	}
	return course, nil
}

func (c *courseClient) SearchCoursesByTitle(title string) (dto.CoursesResponse_Full, error.ApiError) {

	return dto.CoursesResponse_Full{}, nil
}

func (c *courseClient) SearchCoursesByCategory(category string) (dto.CoursesResponse_Full, error.ApiError) {

	return dto.CoursesResponse_Full{}, nil
}

func (c *courseClient) SearchCoursesByDescription(description string) (dto.CoursesResponse_Full, error.ApiError) {

	return dto.CoursesResponse_Full{}, nil
}

func (c *courseClient) CreateCourse(course dto.CourseResponse_Full) (dto.CourseResponse_Full, error.ApiError) {

	return course, nil
}

func (c *courseClient) UpdateCourse(course dto.CourseRequest_Registration) (dto.CourseResponse_Full, error.ApiError) {

	return dto.CourseResponse_Full{}, nil
}

func (c *courseClient) DeleteCourse(courseId int) error.ApiError {

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
