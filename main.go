package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const enHelloPrefix = "Hello, "
const spHelloPrefix = "Hola, "
const frHelloPrefix = "Bonjour, "

func main() {
	fmt.Println(Hello("World", ""))
}

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix := enHelloPrefix

	switch lang {
	case french:
		prefix = frHelloPrefix
	case spanish:
		prefix = spHelloPrefix
	}

	return prefix + name
}
