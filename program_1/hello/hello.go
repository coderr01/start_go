package main

import (
	"fmt"
	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Khushal")
	fmt.Println(message)
}
