package main

import (
	"fmt"
	helloworld "hello-world/hello-world"
)

func main() {
	names := []string{"rick", "bhoomi"}

	messages, err := helloworld.Hellos(names)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(messages)
}
