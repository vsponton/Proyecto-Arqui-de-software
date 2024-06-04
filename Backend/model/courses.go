package model
//model es el dao


type Course struct {
	ID           int64     // Course ID
	Title        string    // Course title
	Description  string    // Course description
	Category     string    // Course Category. Allowed values: to be defined
	ImageURL     string    // Course image URL
	Duration     int64    // Course duration
	Requirements string   //Course requirements
}
