package dto

type NewUserRequest struct {
	Username string `json:"username" validate:"required~username cannot be empty"`
	Email    string `json:"email" validate:"required~email cannot be empty,email,unique"`
	Password string `json:"password" validate:"required~password cannot be empty,min=6,max=72"`
	Age      int    `json:"age" validate:"required~age cannot be empty,numeric"`
}
type NewUserRequestLogin struct {
	Email    string `json:"email" validate:"required~email cannot be empty,email,unique"`
	Password string `json:"password" validate:"required~password cannot be empty,min=6,max=72"`
}

type NewUserResponse struct {
	Result     string `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	Result     string        `json:"result"`
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	Data       TokenResponse `json:"data"`
}
