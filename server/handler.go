package server

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cod3rboy/yaps/img"
	"github.com/cod3rboy/yaps/utils"
	"github.com/cod3rboy/yaps/utils/sliceutils"
	"github.com/cod3rboy/yaps/utils/stringutils"
	"github.com/gofiber/fiber/v2"
)

// Supported image formats
var SupportedFormats = []string{
	img.IMAGE_JPG,
	img.IMAGE_JPEG,
	img.IMAGE_PNG,
	img.IMAGE_TIFF,
	img.IMAGE_WEBP,
}

// Constants for query parameter keys
const (
	keySize      = "s"
	keyBgColor   = "b"
	keyTextColor = "c"
	keyText      = "t"
	keyScale     = "x"
)

// Client Errors
var (
	ErrUnsupportedFormat     = fiber.NewError(fiber.ErrBadRequest.Code, "unsupported image format")
	ErrInvalidParamSize      = fiber.NewError(fiber.ErrBadRequest.Code, "invalid size ("+keySize+") value")
	ErrInvalidParamScale     = fiber.NewError(fiber.ErrBadRequest.Code, "invalid scale ("+keyScale+") value")
	ErrInvalidParamBgColor   = fiber.NewError(fiber.ErrBadRequest.Code, "invalid background color ("+keyBgColor+") value")
	ErrInvalidParamTextColor = fiber.NewError(fiber.ErrBadRequest.Code, "invalid text color ("+keyTextColor+") value")
)

// Constants to help parse size parameter
const (
	dimensionDelimiter = "x"
	widthIndex         = 0
	heightIndex        = 1
)

// Default values for parameters
var (
	defaultSize    = img.Size{Width: 100, Height: 100}
	defaultBgColor = img.Color{
		// Light Gray (#cccccc)
		R: 0xCC,
		G: 0xCC,
		B: 0xCC,
	}
	defaultTextColor = img.Color{
		// Dark Gray (#969696)
		R: 0x96,
		G: 0x96,
		B: 0x96,
	}
	defaultScale = 1.0
)

// HandlerImage is a handler to serve image generation request.
//
// It is only compatible with [fiber.Handler] inteface and not with [http.Handler] interface.
func HandlerImage(ctx *fiber.Ctx) error {
	format := ctx.Params("format", "")
	if !sliceutils.ContainsString(SupportedFormats, format) {
		return ErrUnsupportedFormat
	}
	// Get all image parameters from query string
	size, err := getParamSize(ctx)
	if err != nil {
		return ErrInvalidParamSize
	}
	bgColor, err := getParamBgColor(ctx)
	if err != nil {
		return ErrInvalidParamBgColor
	}
	txtColor, err := getParamTextColor(ctx)
	if err != nil {
		return ErrInvalidParamTextColor
	}
	scale, err := getParamScale(ctx)
	if err != nil {
		return ErrInvalidParamScale
	}
	defaultText := fmt.Sprintf("%d %s %d", utils.ScaleDimension(size.Width, scale), dimensionDelimiter, utils.ScaleDimension(size.Height, scale))

	text := getParamText(ctx, defaultText)

	params := &img.ImageParams{
		Format:          format,
		Size:            size,
		BackgroundColor: bgColor,
		TextColor:       txtColor,
		Scale:           scale,
		Text:            text,
	}

	result, err := img.Generate(params)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	ctx.Set("Content-Type", result.MimeType)
	ctx.Set("Content-Length", string(result.Bytes))

	return ctx.Send(result.Bytes)
}

// getParamSize returns the image size read from query parameters.
//
// If an error occurs, it returns nil, error.
// If no size is present in query parameters, it returns defaultSize.
func getParamSize(ctx *fiber.Ctx) (*img.Size, error) {
	sizeParam := new(img.Size)
	sizeParam.Height = defaultSize.Height
	sizeParam.Width = defaultSize.Width

	sizeValue := ctx.Query(keySize)
	if sizeValue == "" {
		return sizeParam, nil
	}
	sizeValue = strings.ToLower(sizeValue)

	var w, h string

	if !strings.Contains(sizeValue, dimensionDelimiter) {
		// Square side
		w, h = sizeValue, sizeValue
	} else {
		// Rectange width and height
		dimensions := strings.Split(sizeValue, dimensionDelimiter)
		w, h = dimensions[widthIndex], dimensions[heightIndex]
	}

	width, err := strconv.Atoi(w)
	if err != nil {
		return nil, err
	}
	height, err := strconv.Atoi(h)
	if err != nil {
		return nil, err
	}
	sizeParam.Width = width
	sizeParam.Height = height
	return sizeParam, nil
}

// getParamBgColor returns the image background color read from query parameters.
//
// If an error occurs, it return nil, error.
// If no background color is present in query parameters, it returns defaultBgColor.
func getParamBgColor(ctx *fiber.Ctx) (*img.Color, error) {
	bgColorParam := new(img.Color)
	bgColorParam.R = defaultBgColor.R
	bgColorParam.G = defaultBgColor.G
	bgColorParam.B = defaultBgColor.B

	bgColorValue := ctx.Query(keyBgColor)

	if bgColorValue == "" {
		return bgColorParam, nil
	}

	color, err := stringutils.ParseColorHex(bgColorValue)
	if err != nil {
		return nil, err
	}
	red, green, blue := utils.GetRGBComponents(color)

	bgColorParam.R = red
	bgColorParam.G = green
	bgColorParam.B = blue

	return bgColorParam, nil
}

// getParamTextColor returns the image text color read from query parameters.
//
// If an error occurs, it return nil, error.
// If no text color is present in query parameters, it returns defaultTextColor.
func getParamTextColor(ctx *fiber.Ctx) (*img.Color, error) {
	txtColorParam := new(img.Color)
	txtColorParam.R = defaultTextColor.R
	txtColorParam.G = defaultTextColor.G
	txtColorParam.B = defaultTextColor.B

	txtColorValue := ctx.Query(keyTextColor)

	if txtColorValue == "" {
		return txtColorParam, nil
	}

	color, err := stringutils.ParseColorHex(txtColorValue)
	if err != nil {
		return nil, err
	}
	red, green, blue := utils.GetRGBComponents(color)

	txtColorParam.R = red
	txtColorParam.G = green
	txtColorParam.B = blue

	return txtColorParam, nil
}

// getParamText returns the image text read from query parameters.
//
// If no text is present in query parameters, it returns defaultValue.
func getParamText(ctx *fiber.Ctx, defaultValue string) string {
	return ctx.Query(keyText, defaultValue)
}

// getParamScale returns the image scale read from query parameters.
//
// If an error occurs, it returns 0.0, error.
// If no scale is present in query parameters, it returns defaultScale.
func getParamScale(ctx *fiber.Ctx) (float64, error) {
	scaleValue := ctx.Query(keyScale)
	if scaleValue == "" {
		return defaultScale, nil
	}
	scaleParam, err := strconv.ParseFloat(scaleValue, 32)
	if err != nil {
		return 0.0, err
	}
	return float64(scaleParam), nil
}
