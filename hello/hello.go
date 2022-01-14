package main

import (
	"fmt"
	"seacoastcoder.com/greetings"
)

import "rsc.io/quote"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	message := greetings.Hello("Test")
	fmt.Println(message)
}