package main

import "fmt"

func HelloWorld() string {
	messageHelloWorld := "Hello World from Go!"

	//Do not remove this line here
	return messageHelloWorld
}

func main() {

	fmt.Println(HelloWorld())
}
