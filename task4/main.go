// 4. Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и
// выводят в stdout. Необходима возможность выбора количества воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C.
// Выбрать и обосновать способ завершения работы всех воркеров.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func Worker(ch chan interface{}, ctx context.Context, num int, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Stop Worker%d\n", num)
			wg.Done()
			return
		case data := <-ch:
			fmt.Printf("Worker%d: %v\n", num, data)
			time.Sleep(time.Millisecond * 200)
		}
	}
}

func main() {
	mainChan := make(chan interface{})
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	// Args amount check
	var n int
	if len(os.Args) == 2 {
		var err error
		n, err = strconv.Atoi(os.Args[1])
		fmt.Println(n)
		if err != nil || n == 0 {
			fmt.Println("Wrong workers amount:", err)
			return
		}
	} else {
		fmt.Println("Wrong args amount.")
		return
	}
	fmt.Println(n)
	for i := 0; i < n; i++ {

		fmt.Printf("Start Worker%d\n", i)
		wg.Add(1)
		go Worker(mainChan, ctx, i, &wg)
	}

	// Wait for a SIGINT
	// Stop workers using context when signal is received
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	var id int
	fmt.Println("Start posting.")
	go func() {
		for {
			select {
			case signal := <-signalChan:
				fmt.Printf("\nReceived an %s, closing workers\n\n", signal)
				cancel()
				close(mainChan)
				return
			default:
				data := []interface{}{1, 1.1, "Some string", "Слово"}
				for _, el := range data {
					mainChan <- el
				}
				id++
				time.Sleep(time.Millisecond * 250)
			}
		}
	}()
	wg.Wait()
}
