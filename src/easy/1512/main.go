package main

import "fmt"

func numIdenticalPairs(nums []int) int {
    total := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] { total++ }
		}
	}

	return total
}

// func numIdenticalPairs(nums []int) int {
//     result := make(map[int]int)
// 	total := 0

// 	for _, num := range nums {
// 		total += result[num]
// 		result[num]++
// 	}

// 	return total
// }

func main() {
	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
	fmt.Println(numIdenticalPairs([]int{1, 1, 1, 1}))
	fmt.Println(numIdenticalPairs([]int{1, 2, 3}))
}
