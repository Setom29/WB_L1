// 13. Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	a := 1
	b := 2
	fmt.Printf("First number: %d\nSecond number: %d\n", a, b)
	fmt.Println()
	a, b = b, a
	fmt.Printf("First number: %d\nSecond number: %d\n", a, b)
	fmt.Println()
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("First number: %d\nSecond number: %d\n", a, b)
}
