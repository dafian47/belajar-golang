package main

import (
	"fmt"
	"strings"
)

type Scholar struct {
	Name	string
	Age		int8
}

/**
Simple Method
 */
func (s Scholar) sayHello()  {
	fmt.Println("Hello", s.Name)
}

/**
Method with Return-Value
 */
func (s Scholar) getNameAt(i int8) string {
	return strings.Split(s.Name, " ")[i - 1]
}

/**
Method Pointer
 */
func (s *Scholar) sayHai()  {
	fmt.Println("Hai", s.Name)
}

func main() {

	// Use Method
	var s1 = Scholar{"Mohamad Rizki Dafianto", 23}
	s1.sayHello()

	var name = s1.getNameAt(2)
	fmt.Println(name)

	var s2 = &Scholar{"Fahry Saifullah", 21}
	s2.sayHai()
	s2.sayHello()
}
