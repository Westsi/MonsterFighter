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
	fmt.Println("GOT NEW USER REQ")
	// POST
	var parsedUser struct {
		Name     string
		Password string
		Email    string
	}

	err := json.NewDecoder(r.Body).Decode(&parsedUser)
	if err != nil {
		fmt.Println("PANIC")
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
	// GET
	var users []User
	db.Find(&users)
	d, err := json.Marshal(users)
	if err != nil {
		fmt.Println("responding with 500")
		w.WriteHeader(500)
		w.Header().Add("error", "error in marshalling json data. please report this as a bug.")
		return
	}
	w.Write(d)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	// GET
	// will have been passed through from request in headers["username"]
	username := r.Header.Get("username")
	if username == "" {
		w.WriteHeader(400)
		w.Header().Add("error", "username is empty")
		return
	}

	var foundUser User
	db.Where(&User{Name: username}).First(&foundUser)
	if foundUser.Name == "" {
		w.WriteHeader(204)
		w.Header().Add("error", "user does not exist")
		return
	}

	d, err := json.Marshal(foundUser)

	if err != nil {
		w.WriteHeader(500)
		w.Header().Add("error", "error in marshalling json data. please report this as a bug.")
		return
	}
	w.Write(d)
}
