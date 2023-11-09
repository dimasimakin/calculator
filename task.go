package main

import "fmt"

func main() {

	var operator string
	var number1 int8
	var number2 int8

	fmt.Print("expression")
	fmt.Scanln(&number1, &operator, &number2)

	switch operator {

	case "+":
		fmt.Printf("%d %s %d = %d", number1, operator, number2, number1+number2)

	case "-":
		fmt.Printf("%d %s %d = %d", number1, operator, number2, number1-number2)

	case "*":
		fmt.Printf("%d %s %d = %d", number1, operator, number2, number1*number2)

	case "/":
		if number2 == 0.0 {
			fmt.Println("Divide by Zero Situation")
		} else {
			fmt.Printf("%d %s %d = %d", number1, operator, number2, number1/number2)
		}
	default:
		fmt.Println("Invalid Operator")
	}
}
