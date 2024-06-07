package model

//model es el dao

type Course struct {
	ID           int    `gorm:"primaryKey"`                 // Course ID
	Title        string `gorm:"type:varchar(255);not null"` // Course title
	Description  string `gorm:"type:varchar(255);not null"` // Course description
	Category     string `gorm:"type:varchar(255);not null"` // Course Category. Allowed values: to be defined
	ImageURL     string `gorm:"type:varchar(255);not null"` // Course image URL
	Duration     int    `gorm:"type:varchar(255);not null"` // Course duration
	Requirements string `gorm:"type:varchar(255);not null"` //Course requirements
}

type Courses []Course
