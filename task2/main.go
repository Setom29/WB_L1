// Написать программу, которая конкурентно рассчитает значение
// квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"sync"
)

// get squares using sync lib
func SquaresWithSync(arr []int) []int {
	var waitGroup sync.WaitGroup
	for i := range arr {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			arr[i] = arr[i] * arr[i]
		}(i)
	}
	waitGroup.Wait()
	return arr
}

// get squares using chan
func SquaresWithChan(c chan int, arr []int) {
	for i := range arr {
		c <- arr[i] * arr[i]
	}
	close(c)
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	c := make(chan int)
	go SquaresWithChan(c, arr)
	for val := range c {
		fmt.Println("Square value chan = ", val)
	}
	for _, el := range SquaresWithSync(arr) {
		fmt.Println("Square value sync = ", el)
	}
}
