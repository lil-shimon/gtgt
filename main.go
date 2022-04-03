package main

import "fmt"

const enHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("World"))
}

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return enHelloPrefix + name
}
