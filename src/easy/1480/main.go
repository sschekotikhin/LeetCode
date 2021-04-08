package main

import "fmt"

func runningSum(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	storage := nums[0]
	for i := 1; i < len(nums); i++ {
		storage += nums[i]
		nums[i] = storage
	}

	return nums
}

// вариант получше
func another(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i - 1]
	}

	return nums
}

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))
	fmt.Println(runningSum([]int{1, 1, 1, 1, 1}))
	fmt.Println(runningSum([]int{3, 1, 2, 10, 1}))
}
