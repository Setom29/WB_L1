// 11. Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

func Intersection(set1 map[int]bool, set2 map[int]bool) map[int]bool {
	res := make(map[int]bool, len(set1))
	for key := range set1 {
		if set2[key] {
			res[key] = true
		}
	}
	return res

}

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7}
	arr2 := []int{5, 6, 7, 8, 9, 10, 11}

	set1 := make(map[int]bool, len(arr1))
	set2 := make(map[int]bool, len(arr2))
	// fill the first set
	for _, el := range arr1 {
		set1[el] = true
	}
	// fill the second set
	for _, el := range arr2 {
		set2[el] = true
	}
	// show the intersection
	fmt.Println(Intersection(set1, set2))

}
