package main

import "fmt"

// Struct
type gasEngine struct {
	mpg     uint32
	gallons uint32
}

type electricEngine struct {
	mpkwh uint32
	kwh   uint32
}

// Method
func (e electricEngine) mileLeft() uint32 {
	return e.kwh * e.mpkwh
}

func (e gasEngine) mileLeft() uint32 {
	return e.gallons * e.mpg
}

// Interface
type engine interface {
	mileLeft() uint32
}

func canMakeIt(e engine, miles uint32) {
	if miles <= e.mileLeft() {
		fmt.Println("You can make it!")
	} else {
		fmt.Println("You can't make it!")
	}
}

// func intDivision(numerator, denominator int) (int, int, error) {
// 	var err error
// 	if denominator == 0 {
// 		err = errors.New("denominator cannot be zero")
// 		return 0, 0, err
// 	}

// 	var result = numerator / denominator
// 	var remainder = numerator % denominator
// 	return result, remainder, err
// }

// func modify(x *int) {
// 	*x = *x + 5
// }

func main() {
	// numerator := 112
	// denominator := 0

	// var result, remainder, err = intDivision(numerator, denominator)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Printf("Result: %d, Remainder: %d\n", result, remainder)

	// myMap := make(map[string]int)

	// myMap["one"] = 1
	// myMap["two"] = 2
	// myMap["three"] = 3

	// for clm, i := range myMap {
	// 	fmt.Println(clm, "fsd", i)
	// }

	var myEngine gasEngine = gasEngine{30, 10}
	// var myEngine2 = struct {
	// 	mpg     uint8
	// 	gallons uint8
	// }{25, 15}

	// fmt.Println(myEngine.mileLeft())
	// fmt.Println(30 * 10)

	canMakeIt(myEngine, 50)
}
