// 20. Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	res := strings.Split(s, " ")
	for i, j := 0, len(res)-1; i < len(res)/2; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return strings.Join(res, " ")
}

func main() {
	str := "snow dog sun"
	fmt.Println(reverseWords(str))
}
