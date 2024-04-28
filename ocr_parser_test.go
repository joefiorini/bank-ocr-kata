package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNumber(t *testing.T) {
	ocr_number := strings.Trim(`
    _  _     _  _  _  _  _ 
  | _| _||_||_ |_   ||_||_|
  ||_  _|  | _||_|  ||_| _|
`, "\n")
	got := OcrParse(ocr_number)
	assert.Equal(t, "123456789", got)
}

func TestGetNumber2(t *testing.T) {
	ocr_number := strings.Trim(`
 _  _  _  _  _  _  _  _  _ 
| || || || || || || || || |
|_||_||_||_||_||_||_||_||_|

`, "\n")
	got := OcrParse(ocr_number)
	assert.Equal(t, "000000000", got)
}
func TestGetNumber3(t *testing.T) {
	ocr_number := strings.Trim(`
                           
  |  |  |  |  |  |  |  |  |
  |  |  |  |  |  |  |  |  |

`, "\n")
	got := OcrParse(ocr_number)
	assert.Equal(t, "111111111", got)
}
func TestGetNumber4(t *testing.T) {
	ocr_number := strings.Trim(`
 _  _  _  _  _  _  _  _  _ 
 _| _| _| _| _| _| _| _| _|
|_ |_ |_ |_ |_ |_ |_ |_ |_ 

`, "\n")
	got := OcrParse(ocr_number)
	assert.Equal(t, "222222222", got)
}
func TestGetNumber5(t *testing.T) {
	ocr_number := strings.Trim(`
 _  _  _  _  _  _  _  _  _ 
 _| _| _| _| _| _| _| _| _|
 _| _| _| _| _| _| _| _| _|
`, "\n")
	got := OcrParse(ocr_number)
	assert.Equal(t, "333333333", got)
}
func TestGetNumber6(t *testing.T) {
	ocr_number := strings.Trim(`
                           
|_||_||_||_||_||_||_||_||_|
  |  |  |  |  |  |  |  |  |

`, "\n")
	got := OcrParse(ocr_number)
	assert.Equal(t, "444444444", got)
}
func TestGetNumber7(t *testing.T) {
	ocr_number := strings.Trim(`
 _  _  _  _  _  _  _  _  _ 
|_ |_ |_ |_ |_ |_ |_ |_ |_ 
 _| _| _| _| _| _| _| _| _|
`, "\n")
	got := OcrParse(ocr_number)
	assert.Equal(t, "555555555", got)
}
func TestGetNumber8(t *testing.T) {
	ocr_number := strings.Trim(`
 _  _  _  _  _  _  _  _  _ 
|_ |_ |_ |_ |_ |_ |_ |_ |_ 
|_||_||_||_||_||_||_||_||_|
`, "\n")
	got := OcrParse(ocr_number)
	assert.Equal(t, "666666666", got)
}
func TestGetNumber9(t *testing.T) {
	ocr_number := strings.Trim(`
 _  _  _  _  _  _  _  _  _ 
  |  |  |  |  |  |  |  |  |
  |  |  |  |  |  |  |  |  |
`, "\n")
	got := OcrParse(ocr_number)
	assert.Equal(t, "777777777", got)
}
func TestGetNumber10(t *testing.T) {
	ocr_number := strings.Trim(`
 _  _  _  _  _  _  _  _  _ 
|_||_||_||_||_||_||_||_||_|
|_||_||_||_||_||_||_||_||_|
`, "\n")
	got := OcrParse(ocr_number)
	assert.Equal(t, "888888888", got)
}
func TestGetNumber11(t *testing.T) {
	ocr_number := strings.Trim(`
 _  _  _  _  _  _  _  _  _ 
|_||_||_||_||_||_||_||_||_|
 _| _| _| _| _| _| _| _| _|
`, "\n")
	got := OcrParse(ocr_number)
	assert.Equal(t, "999999999", got)
}
