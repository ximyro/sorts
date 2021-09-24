package main

import "fmt"

func mergeSort(arr []int) {
	buf := make([]int, len(arr))
	mergeSortImpl(arr, buf, 0, len(arr)-1)
}

func mergeSortImpl(arr []int, buf []int, l, r int) {
	if l < r {
		m := (l + r) / 2
		mergeSortImpl(arr, buf, l, m)
		mergeSortImpl(arr, buf, m+1, r)

		fmt.Printf("Before l: %d, r: %d, m: %d, buf:%v, arr:%v\n", l, r, m, buf, arr)

		k := l
		i := l
		j := m + 1

		for i <= m || j <= r {
			if j > r || (i <= m && arr[i] < arr[j]) {
				buf[k] = arr[i]
				i++
			} else {
				buf[k] = arr[j]
				j++
			}
			k++
		}
		for i := l; i <= r; i++ {
			arr[i] = buf[i]
		}
		fmt.Printf("After l: %d, r: %d, m: %d, buf:%v, arr:%v\n\n\n", l, r, m, buf, arr)
	}
}

func main() {
	arr := []int{7, 4, 9, 2, 1}
	mergeSort(arr)
	fmt.Printf("%v\n", arr)
}

// 0,
//
//
//
//
