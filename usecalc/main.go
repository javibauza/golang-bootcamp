package main

import (
	"fmt"
	"os"
	"strconv"

	calc "simplecalc"
)

func main() {
	var err error

	num1, _ := strconv.Atoi(os.Args[1])
	operator := os.Args[2]
	num2, _ := strconv.Atoi(os.Args[3])

	var res interface{}
	switch operator {
	case "+":
		res = calc.Addition(num1, num2)
	case "-":
		res = calc.Subtraction(num1, num2)
	case "m":
		fmt.Println("entre")
		res = calc.Multiplication(num1, num2)
	case "/":
		if res, err = calc.Division(num1, num2); err != nil {

		}
	default:
		fmt.Println("Invalid operator")
	}

	fmt.Printf("Result: %v", res)
}
