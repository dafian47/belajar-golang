package main

import (
	"strings"
	"fmt"
	"math/rand"
	"time"
	"math"
)

func main() {

	rand.Seed(time.Now().Unix())

	/**
	Function
	To make code is readable, modular & DRY (Don't Repeat Yourself)
	 */

	var names = []string{"Mohamad", "Rizki", "Dafianto"}
	printMessage("Hello", names)

	random1 := randomWithRange(2, 3)
	fmt.Println(random1)

	random2 := randomWithRange(4, 7)
	fmt.Println(random2)

	divideNumber(3, 2)
	divideNumber(0, 2)
	divideNumber(1, 0)

	var diameter float64 = 2.5
	area, circumference := calculate(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)

	var avg = calculateAverage(1, 2, 3, 2, 1, 5, 3, 2)
	fmt.Println(avg)

	var more = []int{1, 2, 3, 2, 1, 3, 2, 3, 1, 3, 2, 2}
	var avg2 = calculateAverage(more...)
	fmt.Printf("%.2f \n", avg2)

	yourHobbies("Dafian", "Playing Guitar", "Reading Book")

	var hobbies = []string{"Listening Music", "Sleep", "Singing"}
	yourHobbies("Rizki", hobbies...)
}

/**
Function Without No Return Value
 */
func printMessage(message string, arr []string)  {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

/**
Function With Return Value
 */
func randomWithRange(min, max int) int  {
	var value = rand.Int() % (max - min + 1) + min
	return value
}

/**
Function with return
 */
func divideNumber(m, n int)  {
	if m == 0 || n == 0 {
		fmt.Println("Can't Divide with Zero Number!")
		return
	}

	res := m / n
	fmt.Println(res)
}

/**
Function with Multi Return Value
 */
func calculate(d float64) (float64, float64)  {

	var area = math.Pi * math.Pow(d / 2, 2)
	var circumference = math.Pi * d

	return area, circumference
}

/**
Function with Variadic Params
 */
func calculateAverage(numbers ...int) float64  {

	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

/**
Function with normal Param & variadic Param
Req variadic Param must at least
 */
func yourHobbies(name string, hobbies ...string)  {

	var hobbiesString = strings.Join(hobbies, ", ")
	fmt.Println("Hobby", name, "is", hobbiesString)
}
