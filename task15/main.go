// 15. К каким негативным последствиям может привести данный фрагмент кода,
// и как это исправить? Приведите корректный пример реализации.

//   ```
//   var justString string
//   func someFunc() {
//    	v := createHugeString(1 << 10)
//     justString = v[:100]
//   }

//   func main() {
//     someFunc()
//   }
//   ```

package main

var justString string

// 1) it is necessary to use an array of runes instead of the string type to get a slice
// 2) checking the length is necessary when trying to get a slice
// (or we will get panic: runtime error: slice bounds out of range [:100] with capacity ...)
func someFunc() {
	v := createHugeString(1 << 10)
	runeArr := []rune(v)
	if len(runeArr) < 100 {
	}
	justString = string(runeArr[:100])

}

func main() {
	someFunc()

}
