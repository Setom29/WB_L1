// 3. Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

package main

import (
	"fmt"
)

func Squares(arr []int) int {
	var sum int
	ch := make(chan int)
	for _, el := range arr {
		go func(el int) {
			ch <- el * el
		}(el)
	}

	for range arr {
		sum += <-ch
	}

	close(ch)
	return sum
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	sum := Squares(arr)
	fmt.Println(sum)
}
