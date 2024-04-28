package main

import (
	"bufio"
	"fmt"
	"strings"
)

var ocr_lookup_data = map[[3]string]int{
	{"   ", "  |", "  |"}: 1,
	{" _ ", " _|", "|_ "}: 2,
	{" _ ", " _|", " _|"}: 3,
	{"   ", "|_|", "  |"}: 4,
	{" _ ", "|_ ", " _|"}: 5,
	{" _ ", "|_ ", "|_|"}: 6,
	{" _ ", "  |", "  |"}: 7,
	{" _ ", "|_|", "|_|"}: 8,
	{" _ ", "|_|", " _|"}: 9,
	{" _ ", "| |", "|_|"}: 0}

func main() {
	test_str := `    _  _     _  _  _  _  _ 
  | _| _||_||_ |_   ||_||_|
  ||_  _|  | _||_|  ||_| _|
`

	// reader := bufio.NewReader(strings.NewReader(test_str))
	// scanner := bufio.NewScanner(reader)

	columns := [][3]string{}

	lines := strings.Split(test_str, "\n")

	line1_scanner := createScanner(lines[0], 3)
	line2_scanner := createScanner(lines[1], 3)
	line3_scanner := createScanner(lines[2], 3)

	fmt.Printf("--%s--\n", lines[0])
	fmt.Printf("--%s--\n", lines[1])
	fmt.Printf("--%s--\n", lines[2])

	for line1_scanner.Scan() && line2_scanner.Scan() && line3_scanner.Scan() {
		column := [3]string{line1_scanner.Text(), line2_scanner.Text(), line3_scanner.Text()}
		columns = append(columns, column)
	}

	for i := 0; i < len(columns); i++ {
		if value, ok := ocr_lookup_data[columns[i]]; ok {
			fmt.Printf("%d", value)
		} else {
			fmt.Printf("%s", "?")
		}

		// for j := 0; j < len(columns[i]); j++ {

		// 	fmt.Printf("%s\n", columns[i][j])
		// }

	}
	fmt.Println()
}

func is1or4(input [3]string) bool {
	if input[0] == "   " {
		return true
	}

	return false
}

func is2or3(input [3]string) bool {
	if input[1] == " _|" {
		return true
	}

	return false
}

func is5or6(input [3]string) bool {
	if input[1] == "|_ " {
		return true
	}

	return false
}

func createScanner(input string, maxChars int) *bufio.Scanner {
	reader := bufio.NewReader(strings.NewReader(input))
	scanner := bufio.NewScanner(reader)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if len(data) >= maxChars {
			return maxChars, data[:maxChars], nil
		}
		return len(data), data, nil
	})
	return scanner
}
