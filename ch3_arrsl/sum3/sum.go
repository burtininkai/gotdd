package main

// Sum returns sum
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers { // _ blank identifier; range out index, value
		sum += number
	}
	return sum
}

// SumAll gives sum of all slices
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
