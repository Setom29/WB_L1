// 9. Разработать конвейер чисел. Даны два канала:
// в первый пишутся числа (x) из массива, во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	var num int
	numChan := make(chan int)
	doubleChan := make(chan int)
	arr := []int{2, 4, 6, 8, 10}
	wg.Add(1)
	go func() {
		for _, el := range arr {
			numChan <- el
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for range arr {
			num = <-numChan
			doubleChan <- num * 2
		}
		wg.Done()
	}()
	for range arr {
		fmt.Println(<-doubleChan)
	}
	wg.Wait()

}
