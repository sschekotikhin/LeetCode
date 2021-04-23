package main

import "fmt"

func countPoints(points [][]int, queries [][]int) (result []int) {
	for _, circle := range queries {
		counter := 0

		for _, point := range points {
			if (point[0] - circle[0]) * (point[0] - circle[0]) +
			   (point[1] - circle[1]) * (point[1] - circle[1]) <= circle[2] * circle[2] {
				counter++
			}
		}

		result = append(result, counter)
	}

	return
}

func main() {
	fmt.Println(countPoints(
		[][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}},
		[][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}},
	))

	fmt.Println(countPoints(
		[][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}},
		[][]int{{1, 2, 2}, {2, 2, 2}, {4, 3, 2}, {4, 3, 3}},
	))
}
