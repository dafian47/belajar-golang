package main

import (
	"runtime"
	"fmt"
)

func main() {

	/**
	Channel
	Created using keyword `make` and `chan`
	 */

	runtime.GOMAXPROCS(2)

	// Create Channel with name chanMessages
	var chanMessages = make(chan string)

	// Closure to send data to channel
	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("Hello %s", who)
		chanMessages <- data // Sending data to channel
	}

	// Different Goroutine
	go sayHelloTo("Mohamad")
	go sayHelloTo("Rizki")
	go sayHelloTo("Dafianto")

	var message1 = <- chanMessages
	fmt.Println(message1)

	var message2 = <- chanMessages
	fmt.Println(message2)

	var message3 = <- chanMessages
	fmt.Println(message3)

	/**
	Buffered Channel
	 */

	messages := make(chan int, 2)

	go func() {
		for {
			i := <- messages
			fmt.Println("Receive Data", i)
		}
	} ()

	for i := 0; i < 5; i++ {
		fmt.Println("Send Data", i)
		messages <- i
	}
}
