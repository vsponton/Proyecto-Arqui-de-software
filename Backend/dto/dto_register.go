package dto

type RegisterRequest struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}
