package main

import "fmt"

func main() {

	/**
	Arithmetic Operator
	 */
	value := (((2 + 6) % 3) * 4 - 2) / 3
	fmt.Printf("Value = %d \n", value)

	/**
	Comparison Operator
	 */
	var equal = (((2 + 6) % 3) * 4 - 2) / 3
	var isEqual = equal == 2

	fmt.Printf("nilai %d (%t) \n", equal, isEqual)

	/**
	Logical Operator
	 */
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
}