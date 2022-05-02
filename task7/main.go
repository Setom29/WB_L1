// 7. Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
)

func main() {
	c := map[int]int{}
	wg := sync.WaitGroup{}
	n := 100
	mu := sync.Mutex{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			mu.Lock()
			c[i] = i
			mu.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println(len(c))
}
