package courses

import (
	"cursos-ucc/dto"
	"net/http"
	"strconv"
)

func GetCourseByIdUser(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id_user"))

	var coursesDto dto.CoursesDto
	coursesDto, err := service.CourseService.GetCourseByIdUser(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, coursesDto)
}

func GetCourseByTitle(c *gin.Context) {

	var title string
	title = c.Param(title)
	var CoursesResponse_Full dto.CoursesResponse_Full

	CoursesResponse_Full, err := service.CourseService.searchByTitle(title)

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

	CoursesResponse_Full, err := service.CourseService.searchByCategoty(category)

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

	CoursesResponse_Full, err := service.CourseService.searchByTitle(description)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, CoursesResponse_Full)
}

func PostCourse(c *gin.Context) {
	var courseDto dto.CourseDto
	err := c.BindJSON(&courseDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	courseDto, er := service.CourseService.PostCourse(courseDto)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, courseDto)
}

func PutCourse(c *gin.Context) {
	var courseDto dto.CourseDto

	err := c.BindJSON(&courseDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	courseDto, er := service.CourseService.PutCourse(courseDto)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, courseDto)

}

/*
func DeleteCourse(c *gin.Context) {
	var courseDto dto.CourseDto

	err := c.BindJSON(&courseDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	courseDto, err := service.CourseService.DeleteCourse(courseDto)

	c.JSON(http.StatusCreated, courseDto)
}
*/

func DeleteCourse(c *gin.Context) {
	idParam := c.Param("id")
	courseID, err := strconv.Atoi(idParam)
	if err != nil {
		log.Error("Invalid course ID: " + idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	err = service.CourseService.DeleteCourse(courseID)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
