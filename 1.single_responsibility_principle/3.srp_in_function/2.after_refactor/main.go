package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Person struct {
	ID    int64
	Name  string
	Phone string
}

func extractUserIDFromRequest(req *http.Request) (int64, error) {
	err := req.ParseForm()
	if err != nil {
		return 0, err
	}

	userID, err := strconv.ParseInt(req.Form.Get("UserID"), 10, 64)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func loadPersonByID(userID int64) (Person, error) {
	row := DB.QueryRow("SELECT * FROM Users WHERE ID = ?", userID)
	person := Person{}
	err := row.Scan(&person.ID, &person.Name, &person.Phone)
	if err != nil {
		return Person{}, err
	}

	return person, nil
}

func outputPerson(resp http.ResponseWriter, person Person) {
	encoder := json.NewEncoder(resp)
	encoder.Encode(person)
}

func loadUserHandler(resp http.ResponseWriter, req *http.Request) {
	userID, err := extractUserIDFromRequest(req)
	if err != nil {
		resp.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	person, err := loadPersonByID(userID)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	outputPerson(resp, person)
}
