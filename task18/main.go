// 18. Реализовать структуру-счетчик, которая будет инкрементироваться
// в конкурентной среде. По завершению программа должна выводить
// итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mu    sync.Mutex
	wg    sync.WaitGroup
}

func NewCounter() *Counter {
	return &Counter{0, sync.Mutex{}, sync.WaitGroup{}}
}

// increase counter by n
func (c *Counter) increaseCounter(n int) {
	for i := 0; i < n; i++ {
		c.mu.Lock()
		c.count += 1
		c.mu.Unlock()
	}
	c.wg.Done()
}

func main() {
	c := NewCounter()
	n := 10
	for i := 0; i < n; i++ {
		c.wg.Add(1)
		go func() {
			c.increaseCounter(n)
		}()
	}
	c.wg.Wait()
	fmt.Println(c.count)
}
