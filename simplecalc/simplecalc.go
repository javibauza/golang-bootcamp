package simplecalc

import (
	"errors"
	"fmt"
)

//add, subtract, multiply and divide

func Addition(addends ...int) int {
	res := 0
	for _, n := range addends {
		res += n
	}
	return res
}

func Subtraction(subtrahends ...int) int {
	res := subtrahends[0]
	for i := 1; i < len(subtrahends); i++ {
		res -= subtrahends[i]
	}
	return res
}

func Multiplication(multiplicands ...int) int {
	res := 1
	for _, n := range multiplicands {
		res *= n
	}
	fmt.Printf("REs: %d", res)
	return res
}

func Division(dividends ...int) (float64, error) {
	res := dividends[0]
	for i := 1; i < len(dividends); i++ {
		if dividends[i] == 0 {
			return 0.0, errors.New("Error dividing by zero")
		}
		res /= dividends[i]
	}
	return float64(res), nil
}
