package main

import "fmt"

const spanish = "Spanish"
const enHelloPrefix = "Hello, "
const spHelloPrefix = "Hola, "

func main() {
	fmt.Println(Hello("World", ""))
}

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	if lang == spanish {
		return spHelloPrefix + name
	}

	return enHelloPrefix + name
}
