package main

import "fmt"

func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		min, index := minElement(arr[i:])
		arr[i+index] = arr[i]
		arr[i] = min
	}
}

func minElement(arr []int) (int, int) {
	m := arr[0]
	index := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < m {
			m = arr[i]
			index = i
		}
	}

	return m, index
}

func main() {
	arr := []int{5, 6, 8, 3, 2, 1}
	selectionSort(arr)
	fmt.Printf("%v\n", arr)
}
