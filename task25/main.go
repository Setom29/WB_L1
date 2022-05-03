// 25. Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

//using time.After
func CustomSleep(t int) {
	<-time.After(time.Second * time.Duration(t))
}

// using loop
func CustomSleep2(t int) {
	start := time.Now()
	for time.Now().Sub(start).Seconds() < float64(t) {
	}
}

func main() {
	fmt.Println(time.Now())
	CustomSleep(5)
	fmt.Println("After sleep using time.After:", time.Now())
	CustomSleep2(5)
	fmt.Println("After sleep using loop:", time.Now())
}
