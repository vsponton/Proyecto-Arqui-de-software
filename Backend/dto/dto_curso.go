package dto
//dto es domain

import "time"

// busqueda de cursos

type CourseRequest_Category struct {
	Category string `json:"category"`
}

type CourseRequest_Title struct {
	Title string `json:"title"`
}

type CourseRequest_Description struct {
	Description string `json:"description"`
}

//detalle del curso
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

//inscripcion en curso

//mis cursos
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
