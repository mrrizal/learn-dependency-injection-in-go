package main

import "fmt"

type Calculator struct {
	data map[string]float64
}

func (c *Calculator) Calculate(path string) error {
	return nil
}

func (c *Calculator) GetData() map[string]float64 {
	return c.data
}

// DefaultPrinter & CSVPrinter otomatis menjadi turunan dari interface Printer
type Printer interface {
	Output(data map[string]float64) error
}

type DefaultPrinter struct{}

func (d *DefaultPrinter) Output(data map[string]float64) error {
	fmt.Println("default printer!")
	return nil
}

type CSVPrinter struct{}

func (c *CSVPrinter) Output(data map[string]float64) error {
	fmt.Println("csv printer!")
	return nil
}

func main() {
	var calculator Calculator
	calculator.data = map[string]float64{"test": 123.0}
	data := calculator.GetData()

	// print data, the think that i need to change just type of printer
	// no need to add additional function/method to Calculate struct
	// var printer DefaultPrinter
	var printer CSVPrinter
	printer.Output(data)
}
