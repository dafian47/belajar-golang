package main

import "fmt"

func main() {

	var myFirstName string = "Mohamad"
	var myMiddleName string
	myMiddleName = "Rizki"
	myLastName := "Dafianto"

	/**
	TODO Format Print

	Use %s for String
	Use %d for Int
	Use %f for Double
	Use %t for Bool
	 */

	fmt.Printf("My Name is %s %s %s \n", myFirstName, myMiddleName, myLastName)

	var one, two, three, four, five string
	one, two, three, four, five = "1", "2", "3", "4", "5"

	fmt.Printf("All Numbers %s, %s, %s, %s, %s \n",
		one, two, three, four, five)

	number, isExists, decimal, text := 1, false, 3.2, "Ok"

	fmt.Printf("All Variable %d, %t, %f, %s \n",
		number, isExists, decimal, text)

	_ = "Variable Nganggur"

	var testVariable = new(string)

	fmt.Println(testVariable)
	fmt.Println(*testVariable)
}
