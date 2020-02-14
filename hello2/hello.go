package main

import "fmt"

const englishHelloPreffix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPreffix + name
}

func main() {
	fmt.Println(Hello("world"))
}
