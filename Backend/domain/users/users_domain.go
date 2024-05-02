package users

type UserData struct{
	User  string `json:"users"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Token string `json:"token"`
}