package dto

//dto es domain

// PANTALLA DE INICIO - HOME

type CourseResponse_Basic struct {
	ID_Course int64  `json:"id"`
	Title     string `json:"title"`
	Category  string `json:"category"`
	ImageURL  string `json:"image_url"`
}

type CoursesResponse_Basic []CourseResponse_Basic // si está en plural son muchos

type CoursesResponse_Full []CourseResponse_Full // si está en plural son muchos

type CoursesRequest_Category []CourseRequest_Category

type CoursesRequest_Title []CourseRequest_Title

type CoursesRequest_Description []CourseRequest_Description

// BUSQUEDA DE CURSOS

type CourseRequest_Category struct {
	Category string `json:"category"`
}

type CourseRequest_Title struct {
	Title string `json:"title"`
}

type CourseRequest_Description struct {
	Description string `json:"description"`
}

// DETALLE DEL CURSO

type CourseResponse_Full struct {
	ID_Course    int64  `json:"id_course"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Category     string `json:"category"`
	ImageURL     string `json:"image_url"`
	Duration     int64  `json:"duration"`
	Instructor   string `json:"instructor"`
	Requirements string `json:"requirements"`
}

// INSCRIPCION EN CURSO

type CourseRequest_Registration struct {
	// sobre el usuario
	Token string `json:"token"`
	// sobre el curso
	ID_Course int64  `json:"id"`
	Title     string `json:"title"`
}
