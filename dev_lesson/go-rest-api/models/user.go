package models

import "gorm.io/gorm"



type User struct {
	gorm.Model
	Name  string `jeson:"name"`
	Email string `json:"email"`
}