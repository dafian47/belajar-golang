package main

import (
	"reflect"
	"fmt"
)

func main() {

	/**
	Reflect
	 */
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	// Checking Variable Type
	fmt.Println("Tipe Variabel", reflectValue.Type())

	// If Variable Type is Int
	if reflectValue.Kind() == reflect.Int {
		// Show Variable Value (By Casting to Int)
		fmt.Println("Nilai Variabel", reflectValue.Int())
	}

	var testData1 = 25.16334
	reflectValue = reflect.ValueOf(testData1)

	fmt.Println("Tipe Variable TestData1", reflectValue.Type())
	fmt.Println("Nilai Variable TestData1", reflectValue.Interface())

	var testData2 = false
	reflectValue = reflect.ValueOf(testData2)

	fmt.Println("Tipe Variable TestData1", reflectValue.Type())
	fmt.Println("Nilai Variable TestData1", reflectValue.Interface())
}
