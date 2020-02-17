package main

import "fmt"

const englishHelloPreffix = "Hello, "

func Hello(name string) string {
	return englishHelloPreffix + name
}

func main() {
	fmt.Println(Hello("world"))
}
