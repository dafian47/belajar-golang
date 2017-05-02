package main

import (
	"fmt"
	"strconv"
	"strings"
	"errors"
)

func main() {

	var input string
	fmt.Println("Input Some Number")
	fmt.Scanln(&input)

	valid, err := validate(input)
	if valid == false {
		// Use panic to show Detail Error Message
		panic(err.Error())
	}

	number, err := strconv.Atoi(input)
	if err == nil {
		fmt.Println("Is Number", number)
	} else {
		fmt.Println("Not Number", input)
		fmt.Println(err.Error())
	}
}

func validate(input string) (status bool, err error) {
	if strings.TrimSpace(input) == "" {
		// Custom Error
		return false, errors.New("Cannot Empty")
	}
	return true, nil
}
