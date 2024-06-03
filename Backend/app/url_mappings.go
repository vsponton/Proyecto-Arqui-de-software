package app

import (
	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Login Mapping
	router.POST("/login", loginController.Login)

	// Register Mapping
	router.POST("/register/", loginController.Register)

	// Courses Mapping
	router.GET("/course/:id_user", courseController.GetCourseByIdUser)
	router.GET("/course/title=:title", courseController.GetCourseByTitle)
	router.GET("/course/category=:category", courseController.GetCourseByCategory)
	router.GET("/course/description=:description", courseController.GetCourseByDescription)

	router.POST("/course", courseController.PostCourse)

	router.PUT("/course/:id_course", courseController.PutCourse)

	router.DELETE("/course", courseController.DeleteCourse)

	// Users Mapping
	router.GET("/user/:id_user", userController.GetUserByIdUser)
	router.GET("/user/email=:email", userController.GetUserByEmail)

	log.Info("Finishing mappings configurations")
}
