package main

import "fmt"

/**
Single Struct
 */
type student struct {
	Name	string
	Grade	int
}

/**
Multiple Struct
 */
type person struct {
	Age		int
	student
}

func main() {

	/**
	Struct
	 Collection of variable or function with specific name
	 */

	var s1 student
	s1.Name = "Dafian"
	s1.Grade = 2

	fmt.Println("Name Student", s1.Name)
	fmt.Println("Grade Student", s1.Grade)

	/**
	Initialize
	 */

	var s2 = student{}
	s2.Name = "Arif"
	s2.Grade = 3

	var s3 = student{"Diki", 4}

	var s4 = student{Name:"Anang"}

	fmt.Println(s2.Name)
	fmt.Println(s3.Name)
	fmt.Println(s4.Name)

	var p1 = person{}
	p1.Name = "Hariz"
	p1.Grade = 3
	p1.Age = 25

	fmt.Println("Grade 1", p1.Grade)
	fmt.Println("Grade 2", p1.student.Grade)

	var p2 = person{25, s1}

	fmt.Println(p2.student.Name)

	/**
	Anonymous Struct
	 */

	var a1 = struct {
		Name	string
		Length	int
	}{}

	a1.Name = "Giraffe"
	a1.Length = 200

	fmt.Println("Name Animal", a1.Name)
	fmt.Println("Length Animal", a1.Length)

	var a2 = struct {
		Name	string
		Length 	int
	}{
		Name:"Zebra",
		Length:230,
	}

	fmt.Println("Name Animal", a2.Name)
	fmt.Println("Length Animal", a2.Length)

	/**
	Combine struct with slice
	 */

	var allStudents = []student {
		{Name:"Jenifer", Grade:25},
		{Name:"Rohaya", Grade:37},
	}

	for _, student := range allStudents {
		fmt.Println("Name Student", student.Name)
	}
}
