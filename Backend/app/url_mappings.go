package app

import (
	courseController "cursos-ucc/controllers/courses"
	loginController "cursos-ucc/controllers/login"
	userController "cursos-ucc/controllers/users"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// MapUrls asigna las rutas a los controladores correspondientes
func MapUrls(router *gin.Engine) {
	// Login Mapping
	router.POST("/login", loginController.Login)

	// Register Mapping
	router.POST("/register", loginController.Register)

	// Courses Mapping
	router.GET("/course", courseController.GetCourses)
	router.GET("/course/:id_user", courseController.GetCourseByIdUser)
	router.GET("/course/title=:title", courseController.GetCourseByTitle)
	router.GET("/course/category=:category", courseController.GetCourseByCategory)
	router.GET("/course/description=:description", courseController.GetCourseByDescription)

	router.POST("/course", courseController.PostCourse)

	router.PUT("/course/:id_course", courseController.PutCourse)

	router.DELETE("/course", courseController.DeleteCourse)

	// Users Mapping
	router.GET("/user/:id_user", userController.GetUserById)
	router.GET("/user/email=:email", userController.GetUserByEmail)

	log.Info("Finishing mappings configurations")
}
