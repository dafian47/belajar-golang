package main

import (
	"fmt"
	"os"
)

func main() {

	/**
	Defer
	Used to end of a statement
	 */
	exampleDefer()

	/**
	Exit
	Use to force-end of program
	 */
	exampleExit()
}

func exampleDefer()  {
	defer fmt.Println("Hallo") // Show Last
	fmt.Println("Good Morning") // Show Beginning
}

func exampleExit()  {
	defer fmt.Println("Good") // Not Show coz Exit
	os.Exit(1)
	fmt.Println("Good Night") // Not Show coz Exit
}
