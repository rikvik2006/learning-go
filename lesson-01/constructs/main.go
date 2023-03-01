package main

import "fmt"

func main() {
	// if constructor
	num1, num2 := 3, 2

	if num2 == 0 {
		fmt.Println("Cannot divite by 0")
		return
	}

	//division ...
	res := num1 / num2
	// Verb digit
	fmt.Printf("%d / %d = %d\n", num1, num2, res)

	// for construct: print numbers from 0 to 9
	for i := 0; i <= 9; i++ {
		fmt.Println(i)
	}
}
