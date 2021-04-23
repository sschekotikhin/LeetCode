package main

import (
	"fmt"
	"sort"
)

func partition(array *[]int, left, right int) int {
	x := (*array)[right]
	less := left

	for i := left; i < right; i++ {
		if (*array)[i] <= x {
			(*array)[i], (*array)[less] = (*array)[less], (*array)[i]
			less++
		}
	}

	(*array)[less], (*array)[right] = (*array)[right], (*array)[less]

	return less
}

func quickSortImplementation(array *[]int, left, right int) {
	if (left < right) {
		q := partition(array, left, right)
		quickSortImplementation(array, left, q - 1)
		quickSortImplementation(array, q + 1, right)		
	}
}

func quickSort(array *[]int) {
	if (len(*array) > 0) {
		quickSortImplementation(array, 0, len(*array) - 1)
	}
}

func maxIceCream(costs []int, coins int) int {
    sort.Ints(costs)

	for i, cost := range costs {
		if coins - cost >= 0 {
			coins -= cost
		} else {
			return i
		}
	}

	return len(costs)
}

func main() {
	fmt.Println(maxIceCream([]int{1, 3, 2, 4, 1}, 7))
	fmt.Println(maxIceCream([]int{10, 6, 8, 7, 7, 8}, 5))
	fmt.Println(maxIceCream([]int{1, 6, 3, 1, 2, 5}, 20))
}
