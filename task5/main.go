// 5. Разработать программу, которая будет последовательно
// отправлять значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	if len(os.Args) == 1 {
		fmt.Println("The argument list is empty, the number of seconds is required.")
		return
	}

	N, _ := strconv.Atoi(os.Args[1])
	ctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(N))
	wg.Add(1)

	go func() {
		var i int
		for {
			select {
			case <-ctx.Done():
				close(ch)
				fmt.Println("Stop posting")
				wg.Done()
				return
			default:
				ch <- i
				i += 1
				time.Sleep(time.Millisecond * 200)
			}
		}
	}()

	wg.Add(1)

	go func() {
		for num := range ch {
			fmt.Println("Received data from channel:", num)
		}
		wg.Done()
	}()

	wg.Wait()
}
