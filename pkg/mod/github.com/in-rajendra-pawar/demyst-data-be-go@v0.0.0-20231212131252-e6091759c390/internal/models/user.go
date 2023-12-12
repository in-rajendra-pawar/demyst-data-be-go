package models

import (
	"strings"
	"time" 
)

// User model
type User struct {
	ID      	int 	   `json:"id" db:"id" redis:"id" validate:"omitempty"`
	FirstName   string     `json:"first_name" db:"first_name" redis:"first_name" validate:"required,lte=30"`
	LastName    string     `json:"last_name" db:"last_name" redis:"last_name" validate:"required,lte=30"`
	Email       string     `json:"email,omitempty" db:"email" redis:"email" validate:"omitempty,lte=60,email"`
	Password    string     `json:"password,omitempty" db:"password" redis:"password" validate:"omitempty,required,gte=6"`
	CreatedAt   time.Time  `json:"created_at,omitempty" db:"created_at" redis:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty" db:"updated_at" redis:"updated_at"`
	LoginAt     time.Time  `json:"login_date" db:"login_at" redis:"login_at"`
}