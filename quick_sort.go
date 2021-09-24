package main

import "fmt"

func partition(arr []int, l, r int) int {
	x := arr[r] // опорный эелемент
	less := l   // последнее значение в  массиве меньше чем x

	for i := l; i < r; i++ {
		if arr[i] <= x {
			v := arr[i]
			arr[i] = arr[less]
			arr[less] = v
			less++
		}
	}

	v := arr[less]
	arr[less] = arr[r]
	arr[r] = v
	return less
}

func quickSort(arr []int) {
	quickSortImpl(arr, 0, len(arr)-1)
}

func quickSortImpl(arr []int, l, r int) {
	if l < r {
		q := partition(arr, l, r)
		quickSortImpl(arr, l, q-1)
		quickSortImpl(arr, q+1, r)
	}
}

func main() {
	arr := []int{7, 5, 3, 1, 2}
	quickSort(arr)
	fmt.Printf("%v\n", arr)
}
