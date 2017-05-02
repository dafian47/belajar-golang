package main

import "fmt"

type citizens struct {
	name 		string
	age 		int
	weight		float32
	haveMarry	bool
	hobbies		[]string
}

func main() {

	var data = citizens{
		name:"Mohamad Rizki Dafianto",
		age:23,
		weight:48.91,
		haveMarry:false,
		hobbies:[]string{"Playing Guitar", "Listening Music"},
	}

	// Layout Format %b
	fmt.Printf("%b\n", data.age)

	// Layout Format %c
	fmt.Printf("%c\n", 1400) // ո
	fmt.Printf("%c\n", 1235) // ӓ

	// Layout Format %d
	fmt.Printf("%d\n", data.age)

	// Layout Format %e or %E
	fmt.Printf("%e\n", data.weight) // 4.891000e+01
	fmt.Printf("%E\n", data.weight) // 4.891000E+01

	// Layout Format %f or %F
	fmt.Printf("%f\n", data.weight)
	fmt.Printf("%.9f\n", data.weight)
	fmt.Printf("%.2f\n", data.weight)
	fmt.Printf("%.f\n", data.weight)

	// Layout Format %g or %G
	fmt.Printf("%g\n", data.weight)
	fmt.Printf("%G\n", data.weight)

	// Layout Format %o
	fmt.Printf("%o\n", data.age)

	// Layout Format %p
	fmt.Printf("%p\n", &data.name)

	// Layout Format %q
	fmt.Printf("%q\n", `" name \ height "`)

	// Layout Format %s
	fmt.Printf("%s\n", data.name)

	// Layout Format %t
	fmt.Printf("%t\n", data.haveMarry)

	// Layout Format %T
	fmt.Printf("%T\n", data.name)
	fmt.Printf("%T\n", data.age)
	fmt.Printf("%T\n", data.weight)
	fmt.Printf("%T\n", data.haveMarry)
	fmt.Printf("%T\n", data.hobbies)

	// Layout Format %v
	fmt.Printf("%v\n", data)
}
