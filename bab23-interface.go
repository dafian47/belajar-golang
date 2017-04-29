package main

import (
	"math"
	"fmt"
)

type calculate interface {
	wide()		float64
	around()	float64
}

type circle struct {
	diameter	float64
}

func (c circle) radian() float64  {
	return c.diameter / 2
}

func (c circle) wide() float64 {
	return math.Pi * math.Pow(c.radian(), 2)
}

func (c circle) around() float64 {
	return math.Pi * c.diameter
}

type square struct {
	side	float64
}

func (s square) wide() float64 {
	return math.Pow(s.side, 2)
}

func (s square) around() float64 {
	return s.side * 4
}

func main() {

	/**
	Basic Interface
	 */
	var model calculate

	model = square{12.25}
	fmt.Println("==== Persegi")
	fmt.Println("Luas", model.wide())
	fmt.Println("Keliling", model.around())

	model = circle{12.25}
	fmt.Println("==== Lingkaran")
	fmt.Println("Luas", model.wide())
	fmt.Println("Keliling", model.around())
	fmt.Println("Jari-jari", model.(circle).radian())

	/**
	Empty Interface
	 */
	var secret interface{}
	secret = "Hallo Guys"
	fmt.Println(secret)

	secret = 21
	fmt.Println(secret)

	secret = 2.10
	fmt.Println(secret)

	var data map[string]interface{}
	data = map[string]interface{}{
		"one": "Agumon",
		"two": 2.1,
		"three": true,
	}

	fmt.Println(data)

	/**
	Interface with Slice & For
	 */
	var fruits = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10},
		[]string{"manggo", "pineapple", "papaya"},
		"orange",
	}

	for _, each := range fruits {
		fmt.Println(each)
	}
}