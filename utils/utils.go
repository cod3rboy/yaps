package utils

import "math"

func GetRGBComponents(color uint64) (red, green, blue uint8) {
	// Extract each color channel value
	blue = uint8(0xFF & color)
	color = color >> 8
	green = uint8(0xFF & color)
	color = color >> 8
	red = uint8(0xFF & color)
	return
}

func ScaleDimension(size int, scale float64) int {
	newWidth := float64(size) * scale
	return int(math.Ceil(newWidth))
}
