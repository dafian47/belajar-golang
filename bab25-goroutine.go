package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string)  {
	for i := 0; i < till ; i++  {
		fmt.Println(i + 1, message)
	}
}

func main() {

	/**
	Goroutine
	 */

	// For determine the number of processor used in program execution
	runtime.GOMAXPROCS(2)

	go print(5, "Hallo") // Making New Goroutine
	print(5, "Ada Apa")

	var input string
	fmt.Scanln(&input)
}
