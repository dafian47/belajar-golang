package main

import (
	"fmt"
	"strings"
)

type FilterCallback func(string string) bool // Alias Scheme Closure

func main() {

	/**
	Basic Closure
	 */

	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{1, 3, 2, 4, 2, 1, 4, 2, 2}
	min, max := getMinMax(numbers)

	fmt.Println("Min", min, "Max", max)

	/**
	Immediately-Invoked Function Expression (IIFE)
	Variable that accommodated is return-value NOT closure
	 */

	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	} (3)

	fmt.Println("original number :", numbers)
	fmt.Println("filtered number :", newNumbers)

	var data = []string{"wick", "jason", "ethan"}
	var dataContainsO = filter(data, func(string string) bool {
		return strings.Contains(string, "o")
	})
	var dataLength5 = filter(data, func(string string) bool {
		return len(string) == 5
	})

	fmt.Println(dataContainsO)
	fmt.Println(dataLength5)
}

/**
Function as Parameter
 */
func filter(data []string, callback FilterCallback) []string  {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
