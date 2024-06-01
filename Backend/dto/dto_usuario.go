package dto

type UserRequest_Email struct {
	Email string `json:"email"`
}

type UserRequest_ID_User struct {
	ID_User int64 `json:"id"`
}

type UserResponse struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Type      string `json:"type"`
}

type UsersResponse []UserResponse