// 16. Реализовать быструю сортировку массива (quicksort)
// встроенными методами языка.

package main

import "fmt"

func customQuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	split := partition(arr)

	customQuickSort(arr[:split])
	customQuickSort(arr[split:])
}

func partition(arr []int) int {
	pivot := arr[len(arr)/2]

	left := 0
	right := len(arr) - 1

	for {
		for ; arr[left] < pivot; left++ {
		}

		for ; arr[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}

		arr[left], arr[right] = arr[right], arr[left]
	}
}

func main() {

	arr := []int{4, 6, 2, 7, 1, 3, 5, 8, 9, 10}
	fmt.Println(arr)
	customQuickSort(arr)
	fmt.Println("After quicksort: ", arr)
}
