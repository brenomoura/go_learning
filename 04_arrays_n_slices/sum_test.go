package main

import "testing"

func TestSum(t *testing.T) {
	// t.Run("5 numbers", func(t *testing.T) {
	// 	nums := []int{1, 2, 3, 4, 5}
	// 	result := Sum(nums)
	// 	expected := 15

	// 	if expected != result {
	// 		t.Errorf("Expected: '%d', Result: '%d', Given: %v", expected, result, nums)
	// 	}
	// })

	t.Run("Any size collection", func(t *testing.T) {
		nums := []int{1, 2, 3}
		result := Sum(nums)
		expected := 6

		if expected != result {
			t.Errorf("Expected: '%d', Result: '%d', Given: %v", expected, result, nums)
		}
	})
}
