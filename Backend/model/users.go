package model

type User struct {
	ID           int    `gorm:"primaryKey"`                 // User ID
	Email        string `gorm:"type:varchar(255);not null"` // User Email
	PasswordHash string `gorm:"type:varchar(255);not null"` // User Password Hash
	Type         string `gorm:"type:varchar(255);not null"` // User Type. Allowed values: admin, normal
	FirstName    string `gorm:"type:varchar(255);not null"` //User Name
	LastName     string `gorm:"type:varchar(255);not null"` //User LastName
}

type Users []User //arreglo de usuarios
