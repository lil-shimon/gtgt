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

	if lang == spanish {
		return spHelloPrefix + name
	}

	if lang == french {
		return frHelloPrefix + name
	}

	return enHelloPrefix + name
}
