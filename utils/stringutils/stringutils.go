// Package stringutils contains common utility functions for strings.
package stringutils

import (
	"errors"
	"strconv"
	"strings"
)

// IsNumber returns true if str is a numeric string otherwise it returns false.
func IsNumber(str string) bool {
	return strings.IndexFunc(str, func(c rune) bool {
		return c < '0' || c > '9'
	}) == -1
}

// ParseColorHex returns color integer value parsed from given colorHexValue string.
// If an error occurs while parsing, it returns 0, error.
//
// colorHexValue must represent color in hexadecimal format.
// e.g. The following formats are correct -
//
// 1. FE0231 (Red=FE, Green=02, Blue=31)
//
// 2. E3D    (Red=EE, Green=33, Blue=DD)
//
// Incorrect formats -
//
// 1. 0xFE0231 (Should not use 0x prefix)
//
// 2. FFFF     (Total digits must be either 3 or 6)
//
// 3. F3FZ32   (Z is not a hexadecimal digit)
func ParseColorHex(colorHexValue string) (uint64, error) {
	hexDigits := len(colorHexValue)
	if hexDigits != 3 && hexDigits != 6 {
		return 0, errors.New("color hex must contain either 3 or 6 digits")
	}

	if hexDigits == 3 {
		// Convert 3 digits hex value to 6 digits
		newValue := ""
		for _, digit := range colorHexValue {
			newValue += string(digit) + string(digit)
		}
		colorHexValue = newValue
	}

	// Parse first 24 bits of hexadecimal color value
	color, err := strconv.ParseUint(colorHexValue, 16, 24)
	if err != nil {
		return 0, err
	}
	return color, nil
}
