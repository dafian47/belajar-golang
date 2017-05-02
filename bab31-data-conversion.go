package main

import (
	"strconv"
	"fmt"
)

func main() {

	// Convert from String to Int
	var str = "124"
	var num, err = strconv.Atoi(str)

	if err == nil {
		fmt.Println(num) // 124
	}

	// Convert from Int to String
	num = 124
	str = strconv.Itoa(num)

	fmt.Println(str) // "124"

	str = "124"
	var num2, err2 = strconv.ParseInt(str, 10, 64)

	if err2 == nil {
		fmt.Println(num2) // 124
	}

	// Convert Int 64 to String
	var num3 = int64(24)
	var str3 = strconv.FormatInt(num3, 8)

	fmt.Println(str3) // 30

	// Convert String to Float
	var str4 = "24.12"
	var num4, err4 = strconv.ParseFloat(str4, 32)

	if err4 == nil {
		fmt.Println(num4) // 24.1200008392334
	}

	// Convert Float 64 to String
	var num5 = float64(24.12)
	var str5 = strconv.FormatFloat(num5, 'f', 6, 64)

	fmt.Println(str5) // 24.120000

	// Convert String to Bool
	var str6 = "true"
	var bul6, err6 = strconv.ParseBool(str6)

	if err6 == nil {
		fmt.Println(bul6) // true
	}

	// Convert Bool to String
	var bul7 = true
	var str7 = strconv.FormatBool(bul7)

	fmt.Println(str7) // true

	// Casting String ↔ Byte
	var text1 = "halo"
	var b = []byte(text1)

	fmt.Printf("%d %d %d %d \n", b[0], b[1], b[2], b[3])
	// 104 97 108 111


}
