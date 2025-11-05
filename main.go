package main

import (
	"fmt"
	"strconv"
)



func main(){
	fmt.Println("Welcome to my simple calculator by GO")
	fmt.Println("Calculation format: value1 operation value2")
	var choice int = 1
	for choice == 1 {
		fmt.Println("Calculation => 1 \nExit => 2")
		fmt.Scanln(&choice)
		if choice == 2 {
			fmt.Println("Goodbye!")
			return
			}else if choice == 1 {
				value1 := askForInput("First")
				fmt.Print("Enter the operation (+, -, *, /): ")
				operation := askForOperation()
				value2 := askForInput("Second")
				result, err := calculate(value1, value2, operation)
				for err != nil {
					fmt.Println("Error",err)
					value2 = askForInput("Second")
					result, err = calculate(value1, value2, operation)
				}
				fmt.Println(" ---------------------------")
				fmt.Printf("| The result of %v %v %v = %v |\n", value1, operation, value2, result)
				fmt.Println(" ---------------------------")
				}else{
				fmt.Println("Invalid choice")
				fmt.Println("Calculation => 1 \nExit => 2")
				choice = 1
			}
		}
	}
	
	func askForInput(inputOrder string) float64 {
		var input string
		fmt.Printf("Enter the %s value: ", inputOrder)
		fmt.Scanln(&input)
		valueAsFloat , err:= convertInputToValue(input)
		for err != nil{
			fmt.Print("Please Enter a valid input as a number: ")
			fmt.Scanln(&input)
			valueAsFloat , err = convertInputToValue(input)
		}
		return valueAsFloat
	}
	
	func askForOperation() string {
	validOperations := map[string]bool{"+": true, "-": true, "*": true, "/": true}
	var operation string
	fmt.Scanln(&operation)
	for !validOperations[operation] {
		fmt.Print("Please Enter a valid operation: ")
		fmt.Scanln(&operation)
	}
	return operation
}
func convertInputToValue(input string) (float64,error) {
	inputAsFloat, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Printf("%v must be a number\n", input)
	}
	return inputAsFloat, err
}