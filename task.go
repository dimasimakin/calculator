package main

import "fmt"

func main() {

	var expression string
	var operator string
	var number1 int8
	var number2 int8

	fmt.Print("Enter expression : ")
	fmt.Scanln(&expression)

	switch operator {

	case "+":
		fmt.Printf("%.3f %s %.3f = %.3f", number1, operator, number2, number1+number2)

	case "-":
		fmt.Printf("%.3f %s %.3f = %.3f", number1, operator, number2, number1-number2)

	case "*":
		fmt.Printf("%.3f %s %.3f = %.3f", number1, operator, number2, number1*number2)

	case "/":
		if number2 == 0.0 {
			fmt.Println("Divide by Zero Situation")
		} else {
			fmt.Printf("%.3f %s %f = %.3f", number1, operator, number2, number1/number2)
		}
	default:
		fmt.Println("Invalid Operator")
	}
}
