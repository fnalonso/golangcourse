package main

import "fmt"

func main() {
	var name string
	fmt.Println("Enter yout name:")
	if _, err := fmt.Scan(&name); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Hello, %s\n", name)

}
