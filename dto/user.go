package dto

// NewUserRequest  is post request body
type NewUserRequest struct {
	Username string `json:"username" validate:"required~username cannot be empty" example:"Charger Iphone"`
	Email    string `json:"email" validate:"required~email cannot be empty,email,unique" example:"Charger Iphone"`
	Password string `json:"password" validate:"required~password cannot be empty,min=6,max=72" example:"Charger Iphone"`
	Age      int    `json:"age" validate:"required~age cannot be empty,numeric" example:"10"`
}

// NewUserRequestLogin  is post request body
type NewUserRequestLogin struct {
	Email    string `json:"email" validate:"required~email cannot be empty,email,unique" example:"Charger Iphone"`
	Password string `json:"password" validate:"required~password cannot be empty,min=6,max=72" example:"Charger Iphone"`
}

// NewUserResponse  is post response body
type NewUserResponse struct {
	Result     string `json:"result" example:"Charger Iphone"`
	StatusCode int    `json:"statusCode" example:"400"`
	Message    string `json:"message" example:"Charger Iphone"`
}

// TokenResponse  is post response body
type TokenResponse struct {
	Token string `json:"token" example:"Charger Iphone"`
}

// LoginResponse  is post response body
type LoginResponse struct {
	Result     string        `json:"result" example:"Charger Iphone"`
	StatusCode int           `json:"statusCode" example:"400"`
	Message    string        `json:"message" example:"Charger Iphone"`
	Data       TokenResponse `json:"data" example:"Charger Iphone"`
}
