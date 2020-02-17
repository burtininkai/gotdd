package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const english = "English"
const englishHelloPreffix = "Hello, "
const spanishHelloPreffix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello function outputs hello mesage for different languages
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	if language == "" {
		language = english
	}

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPreffix
	case english:
		prefix = englishHelloPreffix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "Spanish"))
}
