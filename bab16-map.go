package main

import "fmt"

func main() {

	/**
	Map
	With contains structure key-value
	 */

	var avatars map[int]string
	avatars = map[int]string{} // Initialize

	avatars[0] = "Aang"
	avatars[1] = "Korra"

	fmt.Println("First Avatar", avatars[0])
	fmt.Println("Last Avatar", avatars[3])
	fmt.Println(avatars)

	var powers = map[string]int{}

	powers["pikachu"] = 47
	powers["bulbasur"] = 39

	fmt.Println("Pikachu Powers", powers["pikachu"])
	fmt.Println("Charmender Powers", powers["charmender"])
	fmt.Println(powers)

	var bikes1 = map[int]string{0:"Mio", 1:"Beat", 2:"Vario"}
	var bikes2 = map[int]string{
		0:"Mio",
		1:"Beat",
		2:"Vario",
	}

	fmt.Println(bikes1)
	fmt.Println(bikes2)

	// Kind Initialize Map
	var init1 = map[string]int{}
	var init2 = make(map[string]int)
	var init3 = *new(map[string]int)

	fmt.Println("Init", init1, init2, init3)

	/**
	Looping Map with for-range
	 */
	for key, value := range bikes1 {
		fmt.Println(key, "\t", value)
	}

	/**
	Delete Item in Map
	 */
	fmt.Println(len(bikes1))
	delete(bikes1, 1)
	fmt.Println(len(bikes1))

	/**
	Detect Some Item in Map
	 */

	var value, isExist = bikes1[2]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("No Item")
	}

	/**
	Combine Map with Slice
	 */

	var data = []map[string]string {
		{"name":"Dafian", "address":"Wisma Asri", "age":"23"},
		{"name":"Anang", "address":"Jl Agus Salim", "age":"19"},
	}

	fmt.Println(len(data))
	fmt.Println(data)
}
