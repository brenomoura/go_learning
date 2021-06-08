package main

import (
	"fmt"
	"reflect"
	"testing"
)

// Sum

func TestSum(t *testing.T) {
	t.Run("5 numbers", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		result := Sum(nums)
		expected := 15

		if expected != result {
			t.Errorf("Expected: '%d', Result: '%d', Given: %v", expected, result, nums)
		}
	})

	t.Run("Any size collection", func(t *testing.T) {
		nums := []int{1, 2, 3}
		result := Sum(nums)
		expected := 6

		if expected != result {
			t.Errorf("Expected: '%d', Result: '%d', Given: %v", expected, result, nums)
		}
	})
}

func BenchmarkSum(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Sum(nums)
	}
}

func ExampleSum() {
	nums := []int{1, 2, 3, 4, 5}
	sumResult := Sum(nums)
	fmt.Println(sumResult)
	// Output: 15
}

// Sum All

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %v, Result: %v", expected, result)
	}
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2}, []int{0, 9})
	}
}

func ExampleSumAll() {
	sumAllResult := SumAll([]int{1, 2}, []int{0, 9})

	fmt.Println(sumAllResult)
	// Output: [3 9]
}

// Sum All Tails

func TestSumAllTails(t *testing.T) {
	assertSliceSum := func(t *testing.T, result, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected: %v, Result: %v", expected, result)
		}
	}

	t.Run("Slices sum", func(t *testing.T) {
		result := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		assertSliceSum(t, result, expected)
	})

	t.Run("Sum with empty slice", func(t *testing.T) {
		result := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		assertSliceSum(t, result, expected)
	})
}
