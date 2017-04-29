package main

import "fmt"

func main() {

	/**
	Slice
	 */

	var fruits = []string{"apple", "grape", "banana", "melon"}

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Change "grape" into "pineapple"
	baFruits[0] = "pinnaple"
	fmt.Println()
	fmt.Println("Change Grape into Pineapple")
	fmt.Println()

	fmt.Println(fruits)   // [apple pineapple banana melon]
	fmt.Println(aFruits)  // [apple pineapple banana]
	fmt.Println(bFruits)  // [pineapple banana melon]
	fmt.Println(aaFruits) // [pineapple]
	fmt.Println(baFruits) // [pineapple]

	/**
	Len Function
	To calculate the width of the existing slice
	 */
	fmt.Printf("Lebar Slice fruits %d \n", len(fruits))

	/**
	Cap Function
	To calculate the existing slice capacity
	 */
	fmt.Printf("Kapasitas Slice fruits %d \n", cap(fruits))

	/**
	Append Function
	To append element into slice
	 */
	var cFruits = append(fruits, "Papaya")
	fmt.Println(fruits)
	fmt.Println(cFruits)

	/**
	Copy Function
	To copy data between Slice/Array
	 */

	var aColors = []string{"Red"}
	var bColors = []string{"Blue", "Yellow"}
	var copiedColors = copy(aColors, bColors)

	fmt.Println(aColors)
	fmt.Println(bColors)
	fmt.Println(copiedColors)
}
