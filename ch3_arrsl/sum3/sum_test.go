package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5} //array
		got := Sum(numbers)
		want := 15

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	firstSlice := []int{1, 2}
	secondSlice := []int{0, 9}
	got := SumAll(firstSlice, secondSlice)
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) { // reflect.DeepEqual is not type safe, allowas to compare string to int as example
		t.Errorf("got %v want %v", got, want)
	}
}
