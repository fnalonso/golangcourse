package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4}
	biggest := biggestNumber(slice...)

	fmt.Println(biggest)

}

func biggestNumber(numbers ...int) int {
	var biggestNumber int
	for _, number := range numbers {
		if number > biggestNumber {
			biggestNumber = number
		}
	}
	return biggestNumber
}
