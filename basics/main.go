package main

import (
	"fmt"
)

func main() {
	a, b := swap("hello", "nope")
	fmt.Println(a, b)

	fmt.Println("Return values")
	fmt.Println(split(17))
}
