package main

import "fmt"

const enHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("World", ""))
}

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	if lang == "Spanish" {
		return "Hola, " + name
	}
	return enHelloPrefix + name
}
