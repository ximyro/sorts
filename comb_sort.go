package main

import "fmt"

const factor float64 = 1.247

func combSort(arr []int) {
	var step float64 = float64(len(arr) - 1)
	for step >= 1 {
		istep := int(step)
		for i := 0; i+istep < len(arr); i++ {
			if arr[i] > arr[i+istep] {
				v := arr[i]
				arr[i] = arr[i+istep]
				arr[i+istep] = v
			}
		}
		step /= factor
	}

	for i := 0; i < len(arr)+1; i++ {
		for j := 0; j+1 < len(arr)-i; j++ {
			if arr[j+1] < arr[j] {
				v := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = v
			}
		}
	}
}

func main() {
	arr := []int{5, 8, 4, 1, 2}
	combSort(arr)
	fmt.Printf("%v", arr)
}
