package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPreffix = "Hello, "
const spanishHelloPreffix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello function outputs hello mesage for different languages
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPreffix + name
	}

	if language == french {
		return frenchHelloPrefix + name
	}

	return englishHelloPreffix + name
}

func main() {
	fmt.Println(Hello("world", "Spanish"))
}
