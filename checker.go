package main

import (
	"strconv"
	"strings"
)

type Result int

const (
	valid          = 2
	Valid   Result = valid
	invalid        = 0
	Invalid Result = invalid
	unknown        = 1
	Unknown Result = unknown
)

func Check(input string) Result {
	if strings.Contains(input, "?") {
		return unknown
	}

	if GetChecksum(input) == 0 {
		return valid
	}

	return invalid
}

func GetChecksum(input string) int {
	return checkDigits(input, 0, 0)
}

// Recursively calculate the checksum of the given account number
// by multiplying each digit by its reverse index + 1 and summing
// the products; return the final result mod 11 for the checksum
func checkDigits(number string, idx int, sum int) int {
	if idx < len(number) {
		digit, _ := strconv.Atoi(string(number[idx]))
		// fmt.Printf("%d*%d+", digit, idx+1)
		return checkDigits(number, idx+1, sum+digit*(10-(idx+1)))
	} else {
		// fmt.Println()
		return sum % 11
	}

}
