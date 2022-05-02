// 17. Реализовать бинарный поиск встроенными методами языка.

package main

import "fmt"

func main() {
	var target int
	fmt.Scan(&target)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	fmt.Println(binarySearch(arr, target))
}

func binarySearch(arr []int, target int) int {
	var start, end, mid int
	start, end = 0, len(arr)-1
	mid = (end + start) / 2
	for start <= end {
		if target < arr[mid] {
			end = mid - 1
		} else if target > arr[mid] {
			start = mid + 1
		} else {
			return mid
		}
		mid = (end + start) / 2
	}
	return -1

}
