package stringutils

import (
	"errors"
	"strconv"
	"strings"
)

func IsNumber(str string) bool {
	return strings.IndexFunc(str, func(c rune) bool {
		return c < '0' || c > '9'
	}) == -1
}

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
