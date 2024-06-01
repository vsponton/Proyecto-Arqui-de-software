package model

import "time"

type Subscription struct {
	ID           int64     // Subscription ID
	UserID       int64     // Subscription User ID
	CourseID     int64     // Subscription Course ID
	InscriptionDate time.Time // Subscription inscription date
}
