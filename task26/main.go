// 26. Разработать программу, которая проверяет,
// что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.
//   Например:
//       abcd — true
//       abCdefAaf — false
//       aabcd — false

package main

import "fmt"

func isUnique(s string) bool {
	counter := make(map[rune]bool, len(s))
	for _, el := range s {
		if !counter[el] {
			counter[el] = true
		} else {
			return false
		}
	}
	return true
}

func main() {

	s := "asdf"
	fmt.Println(isUnique(s))
}
