// 14. Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"reflect"
)

func CheckType(in interface{}) string {
	return fmt.Sprintf(reflect.TypeOf(in).String())
}

func main() {
	ch := make(chan int)
	data := []interface{}{1, 1.1, true, ch, struct{}{}, "string"}
	for _, el := range data {
		fmt.Println("val:", el, "type:", CheckType(el))
	}

}
