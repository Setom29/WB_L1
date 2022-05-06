// 6. Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func UsingContext(wg *sync.WaitGroup) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	wg.Add(1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				// using context lib to exit the goroutine
				fmt.Println("Exit goroutine using context")
				wg.Done()
				return
			default:
				fmt.Println(time.Now())
				time.Sleep(time.Second * 1)
			}
		}
	}(ctx)
}

func UsingChan(wg *sync.WaitGroup, ch chan bool) {
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ch:
				// using true value to exit the goroutine
				fmt.Println("Exit goroutine using chan")
				wg.Done()
				return
			default:
				fmt.Println(time.Now())
				time.Sleep(time.Second * 1)
			}
		}
	}()
}

func main() {
	// context exit part
	wg := sync.WaitGroup{}
	UsingContext(&wg)
	time.Sleep(time.Second * 5)

	// chan exit part
	ch := make(chan bool)
	UsingChan(&wg, ch)
	time.Sleep(time.Second * 5)
	ch <- true

	close(ch)
	wg.Wait()
}
