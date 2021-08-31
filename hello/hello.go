package main

import (
	"fmt"
	"greetings"
)

func main() {
	message := greetings.Hello("Saud")
	fmt.Println(message)
}
