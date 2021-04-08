package main

import "fmt"

func maximumWealth(accounts [][]int) int {
    max := 0
	total := 0

	for _, customer := range accounts {
		for _, v := range customer {
			total += v
		}

		if total > max { max = total }
		total = 0
	}

	return max
}

func main() {
	fmt.Println(maximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))
	fmt.Println(maximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}}))
	fmt.Println(maximumWealth([][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}))
}
