package entity

import "time"

type User struct {
	Id        int       `json:"id" validate:"required"`
	Username  int       `json:"username" validate:"required"`
	Email     int       `json:"email" validate:"required,email"`
	Password  int       `json:"password" validate:"required,min=6,max=255"`
	Age       int       `json:"age" validate:"required,numeric"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
