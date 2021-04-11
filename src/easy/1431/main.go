package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maximum := 0
	for _, candy := range candies {
		if candy >= maximum { maximum = candy }
	}

	result := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		result[i] = candies[i] + extraCandies >= maximum
	}

	return result
}

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
	fmt.Println(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1))
	fmt.Println(kidsWithCandies([]int{12, 1, 12}, 10))
}
