package services

import (
	"cursos-ucc/dto"
	"cursos-ucc/model"
	e "cursos-ucc/utils/errors"
	"strconv"

	//"cursos-ucc/db"
	courseclient "cursos-ucc/clients/course"
	registerclient "cursos-ucc/clients/register"

	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

type courseService struct {
	courseClient   courseclient.CourseClientInterface
	registerClient registerclient.RegisterClientInterface
}

type CourseServiceInterface interface {
	GetCourses() (dto.CoursesResponse_Full, e.ApiError)
	GetCourseById(id int) (dto.CourseResponse_Full, e.ApiError)
	GetCoursesByUser(tokenString string) (dto.CoursesResponse_Full, e.ApiError)
	GetAvailableCoursesByUser(tokenString string) (dto.CoursesResponse_Full, e.ApiError)
	SearchCoursesByTitle(title string) (dto.CoursesResponse_Full, e.ApiError)
	SearchCoursesByCategory(category string) (dto.CoursesResponse_Full, e.ApiError)
	SearchCoursesByDescription(description string) (dto.CoursesResponse_Full, e.ApiError)
	CreateCourse(course dto.CourseResponse_Full) (dto.CourseResponse_Full, e.ApiError)
	UpdateCourse(course dto.CourseResponse_Full) (dto.CourseResponse_Full, e.ApiError)
	DeleteCourse(courseId int) e.ApiError
	RegisterUserToCourse(tokenString string, courseId int) (dto.CourseResponse_Registration, e.ApiError)
}

var (
	CourseService CourseServiceInterface
)

func init() {
	CourseService = &courseService{courseclient.CourseClient, registerclient.RegisterClient}
}

// DEVUELVE TODOS LOS CURSOS
// router.GET("/course", courseController.GetCourses)

func (c *courseService) GetCourses() (dto.CoursesResponse_Full, e.ApiError) {

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

func (c *courseService) GetCourseById(id int) (dto.CourseResponse_Full, e.ApiError) {

	var courseDto dto.CourseResponse_Full

	course, err := c.courseClient.GetCourseById(id)

	if err != nil {
		return dto.CourseResponse_Full{}, err
	}

	courseDto.ID_Course = course.ID
	courseDto.Title = course.Title
	courseDto.Description = course.Description
	courseDto.Category = course.Category
	courseDto.ImageURL = course.ImageURL
	courseDto.Duration = course.Duration
	courseDto.Requirements = course.Requirements

	return courseDto, nil
}

func (c *courseService) GetCoursesByUser(tokenString string) (dto.CoursesResponse_Full, e.ApiError) {

	var coursesDto dto.CoursesResponse_Full
	var courses model.Courses
	var _id int
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte("secret_key"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		for key, value := range claims {
			if key == "id" {
				switch v := value.(type) {
				case float64:
					_id = int(v)
				case string:
					_id, err = strconv.Atoi(v)
					if err != nil {
						log.Fatalf("Error converting ID claim to int: %v", err)
					}
				default:
					log.Fatalf("ID claim is of an unexpected type")
				}
			}
		}
	} else {
		return dto.CoursesResponse_Full{}, e.NewBadRequestApiError("Invalid Token")
	}
	log.Debug(_id)
	courses, err = c.courseClient.GetCoursesByUserId(_id)
	if err != nil {
		return nil, e.NewNotFoundApiError("Couldn't find courses by that user")
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

func (c *courseService) GetAvailableCoursesByUser(tokenString string) (dto.CoursesResponse_Full, e.ApiError) {

	var coursesDto dto.CoursesResponse_Full
	var courses model.Courses
	var _id int
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte("secret_key"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		for key, value := range claims {
			if key == "id" {
				switch v := value.(type) {
				case float64:
					_id = int(v)
				case string:
					_id, err = strconv.Atoi(v)
					if err != nil {
						log.Fatalf("Error converting ID claim to int: %v", err)
					}
				default:
					log.Fatalf("ID claim is of an unexpected type")
				}
			}
		}
	} else {
		return dto.CoursesResponse_Full{}, e.NewBadRequestApiError("Invalid Token")
	}
	log.Debug(_id)
	courses, err = c.courseClient.GetAvailableCoursesByUserId(_id)
	if err != nil {
		return nil, e.NewNotFoundApiError("Couldn't find courses by that user")
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

func (c *courseService) SearchCoursesByTitle(title string) (dto.CoursesResponse_Full, e.ApiError) {

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

func (c *courseService) SearchCoursesByCategory(category string) (dto.CoursesResponse_Full, e.ApiError) {

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

func (c *courseService) SearchCoursesByDescription(description string) (dto.CoursesResponse_Full, e.ApiError) {

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

func (c *courseService) CreateCourse(course dto.CourseResponse_Full) (dto.CourseResponse_Full, e.ApiError) {
	result := Db.Create(&course)
	if result.Error != nil {
		return dto.CourseResponse_Full{}, e.NewInternalServerApiError("e creating course", result.Error)
	}
	return dto.CourseResponse_Full{}, nil

}

func (c *courseService) UpdateCourse(course dto.CourseResponse_Full) (dto.CourseResponse_Full, e.ApiError) {

	var existingCourse dto.CourseResponse_Full

	// Verificar si el curso existe
	if err := Db.Where("id = ?", course.ID_Course).First(&existingCourse).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return dto.CourseResponse_Full{}, e.NewNotFoundApiError("course not found")
		}
		return dto.CourseResponse_Full{}, e.NewInternalServerApiError("e finding course", err)
	}

	// Actualizar el curso
	result := Db.Model(&existingCourse).Updates(course)
	if result.Error != nil {
		return dto.CourseResponse_Full{}, e.NewInternalServerApiError("e updating course", result.Error)
	}

	// Retornar el curso actualizado
	return dto.CourseResponse_Full{}, nil
}

func (c *courseService) DeleteCourse(courseId int) e.ApiError {
	var course dto.CourseResponse_Full

	// Verificar si el curso existe
	if err := Db.Where("id = ?", courseId).First(&course).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return e.NewNotFoundApiError("course not found")
		}
		return e.NewInternalServerApiError("e finding course", err)
	}

	// Eliminar el curso
	if err := Db.Delete(&course).Error; err != nil {
		return e.NewInternalServerApiError("e deleting course", err)
	}
	return nil
}

func (c *courseService) RegisterUserToCourse(tokenString string, courseId int) (dto.CourseResponse_Registration, e.ApiError) {
	var register model.Register
	var registerResp dto.CourseResponse_Registration
	var _id int
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte("secret_key"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		for key, value := range claims {
			if key == "id" {
				switch v := value.(type) {
				case float64:
					_id = int(v)
				case string:
					_id, err = strconv.Atoi(v)
					if err != nil {
						log.Fatalf("Error converting ID claim to int: %v", err)
					}
				default:
					log.Fatalf("ID claim is of an unexpected type")
				}
			}
		}
	} else {
		return dto.CourseResponse_Registration{}, e.NewBadRequestApiError("Invalid Token")
	}

	register, err = c.registerClient.RegisterUserToCourse(_id, courseId)
	if err != nil {
		return dto.CourseResponse_Registration{}, e.NewInternalServerApiError("error creating course", err)
	}
	registerResp.ID_Course = register.CourseID
	return registerResp, nil

}
