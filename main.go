package main

import (
	"fmt"
)

func main() {
	test_str := `    _  _     _  _  _  _  _ 
  | _| _||_||_ |_   ||_||_|
  ||_  _|  | _||_|  ||_| _|
`

	fmt.Println(OcrParse(test_str))
	fmt.Println(Check(test_str))
}
