package types

type RegisterUser struct {
	Username string `json:"username"`
	Email    string `json:"Email"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"Email"`
}

func NewUser(registerUser RegisterUser) (User, error) {

	return User{
		Username: registerUser.Username,
		Email:    registerUser.Email,
	}, nil
}

func ValidateEmail(hashedEmail, plainTextPasword string) bool {
	// Todo
	return true
}
