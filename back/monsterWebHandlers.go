package main

import (
	"encoding/json"
	"net/http"
)

func newMonster(w http.ResponseWriter, r *http.Request) {
	
	// POST
	var parents struct {
		POID string
		PTID string
	}

	err := json.NewDecoder(r.Body).Decode(&parents)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createMonster(&User{}, []string{parents.POID, parents.PTID})
}
