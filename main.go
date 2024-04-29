package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 	test_str := `    _  _     _  _  _  _  _
	//   | _| _||_||_ |_   ||_||_|
	//   ||_  _|  | _||_|  ||_| _|
	// `
	file, err := os.ReadFile("data.txt")

	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not read data file")
		os.Exit(1)
	}

	numbers := strings.Split(string(file), "\n\n")

	for index := range numbers {
		accountNumber := OcrParse(numbers[index])
		fmt.Printf("%s ", accountNumber)
		result := Check(accountNumber)
		if result == Unknown {
			fmt.Printf("ILL")
		} else if result == Invalid {
			fmt.Printf("ERR")
		}
		fmt.Println()
		// fmt.Println(Check(accountNumber))
	}
}
