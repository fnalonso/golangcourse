package main

import "fmt"

func main() {
	var inputNumber int

	fmt.Println("Enter a number")
	if _, err := fmt.Scan(&inputNumber); err != nil {
		fmt.Println(err)
	}

	halfValue, even := half(inputNumber)

	fmt.Printf("Half value: %d, Is a even number: %t\n", halfValue, even)

}


func half(number int)(int, bool){
	return number / 2, number % 2 == 0
}