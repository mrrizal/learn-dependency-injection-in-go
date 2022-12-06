package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func loadUserHandler(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	userID, err := strconv.ParseInt(req.Form.Get("UserID"), 10, 64)
	if err != nil {
		resp.WriteHeader(http.StatusPreconditionFailed)
		return
	}
	row := DB.QueryRow("SELECT * FROM Users WHERE ID = ?", userID)
	person := &Person{}
	err = row.Scan(&person.ID, &person.Name, &person.Phone)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	encoder := json.NewEncoder(resp)
	encoder.Encode(person)
}
