package main

import (
	"net/http"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Password     string
	Email        string
	GakZunnCount int
}

func newUser(w http.ResponseWriter, r *http.Request) {
	// add json parsing
	u := User{
		Name: "NAME HERE",
		Password: "PASSWORD HERE",
		Email: "EMAIL HERE",
		GakZunnCount: int(new_user_currency),
	}
	gz_amount += new_user_currency

	db.Create(&u)
}
