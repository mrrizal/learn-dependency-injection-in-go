package main

import (
	"fmt"
	"io"
	"os"
)

type Calculator struct {
	data map[string]float64
}

func (c *Calculator) Calculate(path string) error {
	return nil
}

func (c *Calculator) Output(writer io.Writer) {
	for path, result := range c.data {
		fmt.Fprintf(writer, "%s -> %.1f\n", path, result)
	}
}

func (c *Calculator) OutputCSV(writer io.Writer) {
	for path, result := range c.data {
		fmt.Fprintf(writer, "%s -> %.1f\n", path, result)
	}
}

func main() {
	var calculator Calculator
	calculator.data = map[string]float64{"test": 123.0}
	calculator.Output(os.Stdout)
	calculator.OutputCSV(os.Stdout)
}
