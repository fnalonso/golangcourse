package main

import "fmt"

func main() {

	fmt.Printf("%t\n", (true && false) || (false && true) || !(false && false))
}
