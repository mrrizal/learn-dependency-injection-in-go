package main

import "net/http"

type Person struct{}

func outputCSV(response http.ResponseWriter, person Person) error {
	return nil
}

func outputJSON(response http.ResponseWriter, person Person) error {
	return nil
}

func BuildOutput(response http.ResponseWriter, format string, person Person) {
	var err error
	switch format {
	case "csv":
		err = outputCSV(response, person)
	case "json":
		err = outputJSON(response, person)
	}
	if err != nil {
		// output a server error and quit
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusOK)
}

func main() {

}
