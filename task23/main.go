// 23. Удалить i-ый элемент из слайса.

package main

import "fmt"

func removeElem(arr []int, ind int) []int {
	if ind >= 0 && ind < len(arr) {
		arr = append(arr[:ind], arr[ind+1:]...)
	}
	return arr
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	i := 5
	fmt.Println(removeElem(arr, i))
}
