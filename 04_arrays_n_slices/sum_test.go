package main

import "testing"

func TestSum(t *testing.T) {
	nums := [5]int{1, 2, 3, 4, 5}
	result := Sum(nums)
	expected := 15

	if expected != result {
		t.Errorf("Expected: '%d', Result: '%d', Given: %v", expected, result, nums)
	}
}
