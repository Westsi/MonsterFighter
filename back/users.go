package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Password string
	Email string
	GakZunnCount int
}

