// 19. Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import (
	"fmt"
)

func reverseString1(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func reverseString2(s string) string {
	var result string
	for _, el := range s {
		result = string(el) + result
	}
	return result
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println("Target string: ", str)
	fmt.Println("First: ", reverseString1(str))
	fmt.Println("Second: ", reverseString2(str))
}
