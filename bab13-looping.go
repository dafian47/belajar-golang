package main

import "fmt"

func main() {

	/**
	Looping
	Only for but contains while, foreach, do while, etc
	 */

	// Kind 1
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// Kind 2
	var i = 0
	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	// Kind 3
	var a = 0
	for {
		fmt.Println("Angka", a)
		a++
		if a == 5 {
			break
		}
	}

	// With break & continue
	for i := 1; i <= 10; i++ {
		if i % 2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	// Nested Looping
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	// Use Label
	outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
