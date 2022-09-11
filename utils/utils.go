// Package utils contains common utility functions used across the application.
package utils

import "math"

// GetRGBComponents returns red, green and blue channel values from lowest 24 bits of the given 64-bit integer.
//
// It discards the high-order byte representing alpha value.
//
// e.g. For integer, 0xFF235FED, the extracted components are -
//	red = 0x23
//	green = 0x5F
//	blue = 0xED
// The high-order byte 0xFF is discarded.
func GetRGBComponents(color uint64) (red, green, blue uint8) {
	// Extract each color channel value
	blue = uint8(0xFF & color)
	color = color >> 8
	green = uint8(0xFF & color)
	color = color >> 8
	red = uint8(0xFF & color)
	return
}

// ScaleDimension returns the scaled size rounded up to an integer value.
//
// e.g.
//
// For size = 10 and scale = 2.0, it returns ceil(10 * 2.0) = 20
//
// For size = 15 and scale = 1.5, it returns ceil(15 * 1.5) = 23
func ScaleDimension(size int, scale float64) int {
	newWidth := float64(size) * scale
	return int(math.Ceil(newWidth))
}
