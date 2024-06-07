package model

type Register struct {
	ID       int `gorm:"primaryKey"`
	UserID   int `gorm:"primaryKey"` // Subscription User ID
	CourseID int `gorm:"primaryKey"` // Subscription Course ID
}

type Registers []Register
