package main

import "fmt"

func main() {
	var firstNumber int;
	var secondNumber int;

	fmt.Print("Enter de first number: ")
	if _, err := fmt.Scan(&firstNumber); err != nil {
		fmt.Println("Error!")
	}

	fmt.Print("Enter de second number: ")
	if _, err := fmt.Scan(&secondNumber); err != nil {
		fmt.Println("Error!")
	}

	if secondNumber > firstNumber {
		fmt.Println(secondNumber % firstNumber)
	} else {
		fmt.Println(firstNumber % secondNumber)
	}


}
