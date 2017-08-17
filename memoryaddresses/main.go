package main

import "fmt"

func main(){
	var name string
	fmt.Print("Hi! Say your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello, ", name)
}
