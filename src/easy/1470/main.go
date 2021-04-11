package main

import "fmt"

func shuffle(nums []int, n int) []int {
	result := make([]int, 0)

	for i := 0; i < n; i++ {
		result = append(result, nums[i])
		result = append(result, nums[n + i])
	}

	return result
}

func main() {
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
	fmt.Println(shuffle([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4))
	fmt.Println(shuffle([]int{1, 1, 2, 2}, 2))
}
