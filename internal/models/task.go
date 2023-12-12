package models

import (
	"strings"
	"time" 
)

// Task model
type Task struct {
	ID      	int		   `json:"id" db:"id" redis:"id" validate:"omitempty"`
	UserID    	int		   `json:"user_id" db:"user_id" redis:"user_id" validate:"omitempty"`
	Title   	string     `json:"first_name" db:"first_name" redis:"first_name" validate:"required,lte=30"`
	CreatedAt   time.Time  `json:"created_at,omitempty" db:"created_at" redis:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty" db:"updated_at" redis:"updated_at"`
	UpdatedBy   time.Time  `json:"updated_by" db:"updated_by" redis:"updated_by"`
}