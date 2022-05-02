// 22. Разработать программу, которая перемножает, делит,
// складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	a := big.NewInt(int64(rand.Int63()))
	b := big.NewInt(int64(rand.Int63()))
	res := big.NewInt(0)

	fmt.Println("a =", a, "\nb =", b)
	fmt.Println("a + b = ", res.Add(a, b))
	fmt.Println("a - b = ", res.Sub(a, b))
	fmt.Println("a * b = ", res.Mul(a, b))
	fmt.Println("a / b = ", res.Div(a, b))
}
