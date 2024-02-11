package main

import "fmt"

func quickSort(values []int) {
	processSort(values, 0, len(values)-1)
}

func processSort(values []int, left, right int) {

	if left < right {
		i := left              // left boundary
		pivot := values[right] // pick a pivot

		for j := left; j < right; j++ {
			if values[j] <= pivot {
				values[i], values[j] = values[j], values[i]
				i++
			}
		}

		values[i], values[right] = values[right], values[i]

		processSort(values, left, i-1)
		processSort(values, i+1, right)
	}
}

func main() {

	nums := []int{9, 32, 4, 2, -24, 3}
	quickSort(nums)
	fmt.Println(nums)
}
