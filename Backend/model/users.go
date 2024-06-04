package model

type User struct {
	ID           int64     // User ID
	Email        string    // User Email
	PasswordHash string    // User Password Hash
	Type         string    // User Type. Allowed values: admin, normal
	FirstName    string    //User Name
	LastName     string    //User LastName 
}

type Users []User     //arreglo de usuarios 