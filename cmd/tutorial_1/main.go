package main

import (
	"errors"
	"fmt"
)

func intDivision(numerator, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("denominator cannot be zero")
		return 0, 0, err
	}

	var result = numerator / denominator
	var remainder = numerator % denominator
	return result, remainder, err
}

func modify(x *int) {
	*x = *x + 5
}

func main() {
	// numerator := 112
	// denominator := 0

	// var result, remainder, err = intDivision(numerator, denominator)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Printf("Result: %d, Remainder: %d\n", result, remainder)

	
}
