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
			// send data array element to the first channel
			numChan <- el
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for range arr {
			// get data from the first channel, mult by 2 and send it to the second
			num = <-numChan
			doubleChan <- num * 2
		}
		wg.Done()
	}()
	for range arr {
		// get the data from the second channel
		fmt.Println(<-doubleChan)
	}
	wg.Wait()

}
