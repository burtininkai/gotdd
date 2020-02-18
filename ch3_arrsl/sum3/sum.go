package main

// Sum returns sum
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers { // _ blank identifier; range out index, value
		sum += number
	}
	return sum
}
