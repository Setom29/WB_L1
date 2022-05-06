// 8. Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import (
	"fmt"
	"os"
	"strconv"
)

// change bit in position i
func changeBit(num int64, pos uint64) (int64, error) {
	var mask int64
	mask = 1 << (pos - 1)

	if len(fmt.Sprintf("%b", num)) < len(fmt.Sprintf("%b", mask)) {
		err := fmt.Errorf("In the binary form of a number, there is no digit at the position %d", pos)
		return 0, err
	}
	return num ^ mask, nil

}

func main() {
	var num int64
	var pos uint64
	var err error
	// args validation
	if len(os.Args) == 3 {
		num, err = strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			fmt.Println("Wrong number value")
			return
		}
		pos, err = strconv.ParseUint(os.Args[2], 10, 64)
		if err != nil {
			fmt.Println("Wrong position value")
			return
		}
	} else {
		fmt.Println("Wrong args amount")
		return
	}
	// entered value
	fmt.Printf("Number: %d\nBinary form: %b\n", num, num)

	// changed value
	changedNum, err := changeBit(num, pos)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Number: %d\nBinary form:  %b\n", changedNum, changedNum)
}
