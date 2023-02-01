package main

import (
	"encoding/json"
	"fmt"
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
	var parsedUser struct {
		Name     string
		Password string
		Email    string
	}

	err := json.NewDecoder(r.Body).Decode(&parsedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u := User{
		Name:         parsedUser.Name,
		Password:     parsedUser.Password,
		Email:        parsedUser.Email,
		GakZunnCount: int(new_user_currency),
	}
	gz_amount += new_user_currency

	db.Create(&u)

	w.WriteHeader(200)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := db.Find(&User{})
	d, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(d)
}
