package dto

//dto es domain

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

type CourseRequest_Token struct {
	Token string `json:"token"`
}

// DETALLE DEL CURSO

type CourseResponse_Full struct {
	ID_Course   int    `json:"id_course"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	ImageURL    string `json:"image_url"`
	Duration    string `json:"duration"`
	//Instructor   string `json:"instructor"`
	Requirements string `json:"requirements"`
}

// INSCRIPCION EN CURSO
type CourseRequest_Registration struct {
	Token     string `json:"token"`
	ID_Course int    `json:"course_id"`
}

type CourseResponse_Registration struct {
	ID_Course int `json:"course_id"`
}
