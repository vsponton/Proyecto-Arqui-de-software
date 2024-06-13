package controllers

import (
	"cursos-ucc/dto"
	service "cursos-ucc/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetCourses(c *gin.Context) {

	var coursesDto dto.CoursesResponse_Full
	coursesDto, err := service.CourseService.GetCourses()
	fmt.Println(coursesDto, err)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, coursesDto)
}

func GetCoursesByUser(c *gin.Context) {
	var tokenDto dto.CourseRequest_Token
	_ = c.BindJSON(&tokenDto)

	var coursesDto dto.CoursesResponse_Full
	coursesDto, err := service.CourseService.GetCoursesByUser(tokenDto.Token)

	log.Debug(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	log.Debug(coursesDto)
	c.JSON(http.StatusOK, coursesDto)
}

func GetAvailableCoursesByUser(c *gin.Context) {
	var tokenDto dto.CourseRequest_Token
	_ = c.BindJSON(&tokenDto)

	var coursesDto dto.CoursesResponse_Full
	coursesDto, err := service.CourseService.GetAvailableCoursesByUser(tokenDto.Token)

	log.Debug(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	log.Debug(coursesDto)
	c.JSON(http.StatusOK, coursesDto)
}

func GetCourseById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	var courseDto dto.CourseResponse_Full
	courseDto, err := service.CourseService.GetCourseById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, courseDto)
}

func GetCourseByTitle(c *gin.Context) {

	//coursesDto, err := courseService.GetCourseByIdUser(id) ---> ver si va o no

	var title string
	title = c.Param(title)
	var CoursesResponse_Full dto.CoursesResponse_Full

	CoursesResponse_Full, err := service.CourseService.SearchCoursesByTitle(title)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, CoursesResponse_Full)
}

func GetCourseByCategory(c *gin.Context) {

	var category string
	category = c.Param(category)
	var CoursesResponse_Full dto.CoursesResponse_Full

	CoursesResponse_Full, err := service.CourseService.SearchCoursesByCategory(category)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, CoursesResponse_Full)
}

func GetCourseByDescription(c *gin.Context) {

	var description string
	description = c.Param(description)
	var CoursesResponse_Full dto.CoursesResponse_Full

	CoursesResponse_Full, err := service.CourseService.SearchCoursesByDescription(description)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, CoursesResponse_Full)
}

func PostCourse(c *gin.Context) {
	var courseDto dto.CourseResponse_Full
	err := c.BindJSON(&courseDto)

	if err != nil {
		//log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	courseDto, er := service.CourseService.CreateCourse(courseDto)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, courseDto)
}

func PutCourse(c *gin.Context) {
	var courseDto dto.CourseResponse_Full

	err := c.BindJSON(&courseDto)

	if err != nil {
		//log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	courseDto, er := service.CourseService.UpdateCourse(courseDto)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, courseDto)

}

func DeleteCourse(c *gin.Context) {
	idParam := c.Param("id")
	courseID, err := strconv.Atoi(idParam)
	if err != nil {
		//log.Error("Invalid course ID: " + idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	err = service.CourseService.DeleteCourse(courseID)
	if err != nil {
		//log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func RegisterUserToCourse(c *gin.Context) {
	var crr dto.CourseRequest_Registration
	_ = c.BindJSON(&crr)

	var CourseResponseDto dto.CourseResponse_Registration
	CourseResponseDto, err := service.CourseService.RegisterUserToCourse(crr.Token, crr.ID_Course)
	if err != nil {
		//log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, CourseResponseDto)
}
