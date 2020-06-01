package main

import "fmt"

func binarySearch(a []int, key int) bool {
	start := 0
	end := len(a) - 1
	var i int
	for {
		if end < start {
			return false
		}
		i = (start + end) / 2

		if a[i] == key {
			return true
		}

		if a[i] < key {
			start = i + 1
		} else {
			end = i - 1
		}
	}
}

func main() {
	result := binarySearch([]int{1, 2, 3, 4, 5}, 2)

	if result {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
