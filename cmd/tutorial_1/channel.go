package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sumSlice[T any](slice []T) bool {
	return len(slice) > 0
}

var MAX_CHICKEN_PRICE float32 = 10

func main() {
	// var chickenChannel = make(chan string)
	// var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}

	// for i := range websites {
	// 	go checkChickenPrices(websites[i], chickenChannel)
	// }
	// sendMessage(chickenChannel)

	var intSlice = []int{1, 2, 3, 4, 5}
	fmt.Println(sumSlice[int](intSlice))
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20

		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string) {
	fmt.Printf("\nFound a deal on chicken at %s", <-chickenChannel)
}
