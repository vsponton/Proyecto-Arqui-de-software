package users

import (
	"backend/domain/users"
	usersClient "mvc-go/clients/users"
)

func Login(request users.LoginRequest) users.LoginResponse {
	var user model.User = usersClient.GetUserByUsername(request.Username)
	if user.password!=request.Password {
	return 	users.LoginResponse{
		Token: "-1",
	}

	}
	return users.LoginResponse{
		Token: "1",
	}

	
}

