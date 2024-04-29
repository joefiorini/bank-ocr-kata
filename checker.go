package main

import (
	"strconv"
)

func Check(input string) bool {
	return checkDigits(input, 0, 0)%11 == 0
}

func GetChecksum(input string) int {
	return checkDigits(input, 0, 0)
}

func checkDigits(number string, idx int, sum int) int {
	if idx < len(number) {
		digit, _ := strconv.Atoi(string(number[idx]))
		// fmt.Printf("%d*%d+", digit, idx+1)
		return checkDigits(number, idx+1, sum+digit*(10-(idx+1)))
	} else {
		// fmt.Println()
		return sum
	}

}
