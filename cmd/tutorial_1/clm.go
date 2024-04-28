package main

import "fmt"

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("memory location: %p \n", thing2)
	for i := range *thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}

func main() {
	// var slice = []int32{1, 2, 3, 4}
	// var sliceCopy = slice

	// sliceCopy[0] = 100

	// fmt.Println(slice)
	// fmt.Println(sliceCopy)

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("thing1: %v, memory location: %p \n", thing1, &thing1)

	var thing2 = square(&thing1)
	fmt.Printf("thing2: %v, memory location: %p \n", thing2, &thing2)
	fmt.Printf("thing1: %v, memory location: %p \n", thing1, &thing1)
}
