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

func main() {
	wg := sync.WaitGroup{}
	UsingContext(&wg)
	wg.Wait()
}
