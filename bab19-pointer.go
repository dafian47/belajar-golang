package main

import "fmt"

func main() {

	/**
	Pointer
	Allocation Place Memory
	 */

	var number1 int = 8
	var number2 *int = &number1

	fmt.Println("Number 1 Value", number1) // Result 8
	fmt.Println("Number 1 Address", &number1) // Result 0xc42007c050

	fmt.Println("Number 2 Value", *number2) // Result 8
	fmt.Println("Number 2 Address", number2) // Result 0xc42007c050

	// Change Variable number1 to 10
	// The value of the variable number2 has change too

	number1 = 10

	fmt.Println("Number 1 Value", number1) // Result 10
	fmt.Println("Number 1 Address", &number1) // Result 0xc42007c050

	fmt.Println("Number 2 Value", *number2) // Result 10
	fmt.Println("Number 2 Address", number2) // Result 0xc42007c050

	var num = 3

	fmt.Println(num)

	change(&num, 9)

	fmt.Println(num)
}

/**
	Pointer as Parameter in Function
	 */
func change(num *int, value int)  {
	*num = value
}