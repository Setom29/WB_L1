// 8. Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import (
	"fmt"
	"os"
	"strconv"
)

// changeBit - changes bit in position i
func changeBit(num int64, pos uint64) int64 {
	pos -= 1

	if hasBit(num, pos) {
		num &= ^(1 << pos)
	} else {
		num |= (1 << pos)
	}
	return num
}

// hasBit checks for set or unset bit
func hasBit(num int64, pos uint64) bool {
	val := num & (1 << pos)
	return val > 0
}

func main() {
	var num int64
	var pos uint64
	var err error
	// Args validation
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
	// Entered value
	fmt.Printf("Number: %d\n in bytes: %08b\n", num, num)

	// Changed value
	changedNum := changeBit(num, pos)
	fmt.Printf("Number: %d\n in bytes:  %08b\n", changedNum, changedNum)
}
