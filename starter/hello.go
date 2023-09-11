package main

import (
	"fmt"

	"rsc.io/quote"

	"exp/greeting"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	message := greeting.Hello("Mike")
	fmt.Println(message)
}
