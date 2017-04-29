package library

import "fmt"

type Book struct {
	Title	string
	Year	int
}

/**
Called for First-Time
 */
func init() {

	fmt.Println("--> Called Init Function")
}

/**
Public Function
 */
func SayHelloGuys()  {
	fmt.Println("Hello Guys")
}

/**
Public Function and access to Private Function
 */
func IntroduceSelf(name string)  {
	introduce(name)
}

/**
Private Function
 */
func introduce(name string)  {
	fmt.Println("I'm", name)
}
