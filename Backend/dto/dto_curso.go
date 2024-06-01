package dto
//dto es domain

import "time"

// PANTALLA DE INICIO - HOME

type CourseResponse_Category struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	Category string `json:"category"`
	ImageURL     string    `json:"image_url"`
}

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

type CourseResponse struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Category     string    `json:"category"`
	ImageURL     string    `json:"image_url"`
	CreationDate time.Time `json:"creation_date"`
	Duration     int64     `json:"duration"`
	Instructor   string    `json:"instructor"`
	Requirements string    `json:"requirements"`
}

//DETALLE DEL CURSO

type CourseResponse_Detail struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Category     string    `json:"category"`
	ImageURL     string    `json:"image_url"`
	CreationDate time.Time `json:"creation_date"`
	Duration     int64     `json:"duration"`
	Instructor   string    `json:"instructor"`
	Requirements string    `json:"requirements"`
}

// INSCRIPCION EN CURSO

type CourseRequest_Registration struct {
	// sobre el usuario
	Token string `json:"token"`
	Email    string `json:"email"`
	// sobre el curso
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
}

type CourseResponse_Registration struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Category     string    `json:"category"`
	ImageURL     string    `json:"image_url"`
}


// MIS CURSOS
type CourseResponse struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Category     string    `json:"category"`
	ImageURL     string    `json:"image_url"`
	CreationDate time.Time `json:"creation_date"`
	Duration     int64     `json:"duration"`
	Instructor   string    `json:"instructor"`
	Requirements string    `json:"requirements"`
}
