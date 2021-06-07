package main

func Sum(nums [5]int) int {
	sum := 0
	for num := range nums {
		sum += num
	}

	return sum
}
