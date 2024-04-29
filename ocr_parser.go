package main

import (
	"bufio"
	// "fmt"
	"strconv"
	"strings"
)

var ocrLookupData = map[[3]string]int{
	{
		"   ",
		"  |",
		"  |"}: 1,
	{
		" _ ",
		" _|",
		"|_ "}: 2,
	{
		" _ ",
		" _|",
		" _|"}: 3,
	{
		"   ",
		"|_|",
		"  |"}: 4,
	{
		" _ ",
		"|_ ",
		" _|"}: 5,
	{
		" _ ",
		"|_ ",
		"|_|"}: 6,
	{
		" _ ",
		"  |",
		"  |"}: 7,
	{
		" _ ",
		"|_|",
		"|_|"}: 8,
	{
		" _ ",
		"|_|",
		" _|"}: 9,
	{
		" _ ",
		"| |",
		"|_|"}: 0}

func OcrParse(input string) string {
	columns := [][3]string{}

	lines := strings.Split(input, "\n")

	line1Scanner := createScanner(lines[0], 3)
	line2Scanner := createScanner(lines[1], 3)
	line3Scanner := createScanner(lines[2], 3)

	// fmt.Printf("--%s--\n", lines[0])
	// fmt.Printf("--%s--\n", lines[1])
	// fmt.Printf("--%s--\n", lines[2])

	// Read in the next 3 characters of each line and add them to the columns slice
	for line1Scanner.Scan() && line2Scanner.Scan() && line3Scanner.Scan() {
		column := [3]string{line1Scanner.Text(), line2Scanner.Text(), line3Scanner.Text()}
		columns = append(columns, column)
	}

	var actualNumber strings.Builder

	// Lookup each column in the data map and build a string
	// out of each entry's values
	for i := 0; i < len(columns); i++ {
		if value, ok := ocrLookupData[columns[i]]; ok {
			actualNumber.Write([]byte(strconv.Itoa(value)))
		} else {
			actualNumber.WriteRune('?')
		}

	}

	return actualNumber.String()

}

// Creates a scanner that reads in text 3 characters at a time
// to be used to build an array of columns from the OCR text
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
