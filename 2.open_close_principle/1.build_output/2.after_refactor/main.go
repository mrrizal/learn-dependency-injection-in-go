package main

import "net/http"

type Person struct{}

type PersonFormatter interface {
	Format(http.ResponseWriter, Person) error
}

type CSVFormatter struct{}

func (c *CSVFormatter) Format(resp http.ResponseWriter,
	person Person) error {
	return nil

}

type JSONFormatter struct{}

func (j *JSONFormatter) Format(resp http.ResponseWriter,
	person Person) error {
	return nil
}

type XLSFormatter struct{}

func (x *XLSFormatter) Format(resp http.ResponseWriter,
	person Person) error {
	return nil
}

func BuildOutput(response http.ResponseWriter,
	formatter PersonFormatter, person Person) {
	err := formatter.Format(response, person)
	if err != nil {
		// output a server error and quit
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusOK)
}

func main() {
	var resp http.ResponseWriter
	var person Person
	var csvFormatter CSVFormatter
	var jsonFormatter JSONFormatter
	var xlsFormatter XLSFormatter

	// output to csv
	BuildOutput(resp, &csvFormatter, person)

	// output to json
	BuildOutput(resp, &jsonFormatter, person)

	// output to xls
	BuildOutput(resp, &xlsFormatter, person)
}
