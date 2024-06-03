package users

type userService struct{}

func (s *userService) InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError) {

	var user model.User

	user.FirstName = userDto.FirstName
	user.LastName = userDto.LastName
	user.Username = userDto.Username
	user.Password = userDto.Password
	user.Email = userDto.Email

	user = s.userClient.InsertUser(user)

	userDto.UserId = user.UserId

	return userDto, nil
}

func (s *userService) Login(loginDto dto.LoginDto) (dto.LoginResponseDto, e.ApiError) {

	var user model.User
	user, err := s.userClient.GetUserByUsername(loginDto.Username)
	var loginResponseDto dto.LoginResponseDto
	loginResponseDto.UserId = -1
	if err != nil {
		return loginResponseDto, e.NewBadRequestApiError("Usuario no encontrado")
	}
	if user.Password != loginDto.Password && loginDto.Username != "encrypted" {
		return loginResponseDto, e.NewUnauthorizedApiError("Contraseña incorrecta")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": loginDto.Username,
		"pass":     loginDto.Password,
	})
	var jwtKey = []byte("secret_key")
	tokenString, _ := token.SignedString(jwtKey)
	if user.Password != tokenString && loginDto.Username == "encrypted" {
		return loginResponseDto, e.NewUnauthorizedApiError("Contraseña incorrecta")
	}

	loginResponseDto.UserId = user.UserId
	loginResponseDto.Token = tokenString
	log.Debug(loginResponseDto)
	return loginResponseDto, nil
}
