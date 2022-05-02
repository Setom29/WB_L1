// 12. Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.

package main

import (
	"fmt"
)

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool, len(strings))
	for _, el := range strings {
		set[el] = true
	}

	for key := range set {
		fmt.Print(key, " ")
	}
	fmt.Println()

}
