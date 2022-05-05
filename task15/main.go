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

func someFunc() {
	v := createHugeString(1 << 10)
	runeArr := []rune(v)
	if len(runeArr) < 100 {
		// return error or the full string according to the task
	}
	justString = string(runeArr[:100])

}

func main() {
	someFunc()

}
