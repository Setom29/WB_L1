// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

package main

import (
	"fmt"
	_ "os"
)

type Human struct {
	Age  int64
	Name string
}

type Action struct {
	Human
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, my name is %s, i'm %d.\n", h.Name, h.Age)
}

func main() {
	human := Human{Name: "John", Age: 35}
	action := Action{Human: human}
	action.SayHi()
}
