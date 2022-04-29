// 3. Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

package main

import (
	"fmt"
	"sync"
)

func SquaresWithSync(arr []int) int {
	var waitGroup sync.WaitGroup
	var sum int
	var mu sync.Mutex
	for i := range arr {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()

			sum += arr[i] * arr[i]
		}(i)

	}
	waitGroup.Wait()
	return sum
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	sum := SquaresWithSync(arr)
	fmt.Println(sum)
}
