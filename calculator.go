package main

import (
	"errors"
)


func calculate(input1 float64, input2 float64, operation string) (float64, error) {
	var result float64
	var err error = nil
	switch operation {
	case "+":
		result = addValues(input1, input2)
	case "-":
		result = subtractValues(input1, input2)
	case "*":
		result = multiplyValues(input1, input2)
	case "/":
		result,err = divideValues(input1, input2)
	}
	return result, err
}




func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) (float64, error) {
	var err error = nil
	if value2 == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return value1 / value2, err
}
