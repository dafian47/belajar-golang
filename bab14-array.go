package main

import "fmt"

func main() {

	/**
	Array

	Start index from 0
	 */

	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])

	/**
	With Horizontal & Vertical Way
	 */

	var fruits [4]string

	// With Horizontal
	fruits  = [4]string{"apple", "grape", "banana", "melon"}

	// With Vertical
	fruits  = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	/**
	Without Element Count
	 */
	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	/**
	Array Multi-Dimension
	 */
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	/**
	Lopping Array
	 */
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}

	/**
	Looping Array with for-range
	 */
	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	/**
	Allocation Array with make
	 */
	var colors = make([]string, 2)
	colors[0] = "Red"
	colors[1] = "Blue"

	fmt.Println(colors)

	var bindings = make([]string, 4)
	bindings[0] = "Fire"
	bindings[1] = "Water"
	bindings[2] = "Earth"
	bindings[3] = "Air"

	fmt.Println(bindings)
}
