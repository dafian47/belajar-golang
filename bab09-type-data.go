package main

import "fmt"

func main() {

	/**
	Type Data Numeric
	 */

	/**
	For Type Non-Decimal Numeric

	Start with uint8, uint16, uint32, uint64 for Positive Number
	Start with int8, int16, int32, int64 for Negative Number

	byte equals to uint8
	uint equals to uint32 or uint64
	int equals to int32 or int64
	rune equals to int32

	 */

	var positiveNumber uint8 = 89
	var negativeNumber = -123456789

	fmt.Printf("Positive Number %d \n", positiveNumber)
	fmt.Printf("Negative Number %d \n", negativeNumber)

	/**
	For Type Decimal Numeric

	Only float32 and float64

	 */

	var myDecimalNumber = 2.64

	fmt.Printf("My Decimal Number %f \n", myDecimalNumber)
	fmt.Printf("My Decimal Number %.2f \n", myDecimalNumber)

	// %f will result 2.640000
	// %.2f will result 2.64
	// %.nf will result 2.n

	/**
	Type Data Bool

	Only true or false

	 */

	var isDeleted bool = false
	fmt.Printf("Is Deleted %t \n", isDeleted)

	isValid := true
	fmt.Printf("IS Valid %t \n", isValid)

	/**
	Type Data String
	 */

	var message string = "This Message"
	var longMessage string = `This Is Long Message
	that can't escape
	Awesome, isn't? `

	fmt.Printf("Message %s \n", message)
	fmt.Printf("Long Message %s \n", longMessage)

	/**
	Default Value Type Data
	- String -> ""
	- Numeric Non-Decimal -> 0
	- Numeric Decimal -> 0.0
	- Bool -> false
	 */
}
