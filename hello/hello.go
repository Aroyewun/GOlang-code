package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	//Get a greeting message and print it.
	name := "Anuoluwapo"
	age := 25
	message := greetings.Hello(name, age)

	fmt.Println(message)
}
