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

// Query string keys for parameters
const KEY_FOR_SIZE = "s"
const KEY_FOR_BG_COLOR = "b"
const KEY_FOR_TEXT_COLOR = "c"
const KEY_FOR_TEXT = "t"
const KEY_FOR_SCALE = "x"

// Client Errors
var ErrUnsupportedFormat = fiber.NewError(fiber.ErrBadRequest.Code, "unsupported image format")
var ErrInvalidParamSize = fiber.NewError(fiber.ErrBadRequest.Code, "invalid size ("+KEY_FOR_SIZE+") value")
var ErrInvalidParamScale = fiber.NewError(fiber.ErrBadRequest.Code, "invalid scale ("+KEY_FOR_SCALE+") value")
var ErrInvalidParamBgColor = fiber.NewError(fiber.ErrBadRequest.Code, "invalid background color ("+KEY_FOR_BG_COLOR+") value")
var ErrInvalidParamTextColor = fiber.NewError(fiber.ErrBadRequest.Code, "invalid text color ("+KEY_FOR_TEXT_COLOR+") value")

// Constants to help parse size parameter
const DIMENSION_DELIMITER = "x"
const WIDTH_INDEX = 0
const HEIGHT_INDEX = 1

// Default values for parameters
var DEFAULT_SIZE = img.Size{Width: 100, Height: 100}
var DEFAULT_BG_COLOR = img.Color{
	// Light Gray (#cccccc)
	R: 0xCC,
	G: 0xCC,
	B: 0xCC,
}
var DEFAULT_TEXT_COLOR = img.Color{
	// Dark Gray (#969696)
	R: 0x96,
	G: 0x96,
	B: 0x96,
}
var DEFAULT_SCALE = 1.0

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
	defaultText := fmt.Sprintf("%d %s %d", utils.ScaleDimension(size.Width, scale), DIMENSION_DELIMITER, utils.ScaleDimension(size.Height, scale))

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

func getParamSize(ctx *fiber.Ctx) (*img.Size, error) {
	sizeParam := new(img.Size)
	sizeParam.Height = DEFAULT_SIZE.Height
	sizeParam.Width = DEFAULT_SIZE.Width

	sizeValue := ctx.Query(KEY_FOR_SIZE)
	if sizeValue == "" {
		return sizeParam, nil
	}
	sizeValue = strings.ToLower(sizeValue)

	var w, h string

	if !strings.Contains(sizeValue, DIMENSION_DELIMITER) {
		// Square side
		w, h = sizeValue, sizeValue
	} else {
		// Rectange width and height
		dimensions := strings.Split(sizeValue, DIMENSION_DELIMITER)
		w, h = dimensions[WIDTH_INDEX], dimensions[HEIGHT_INDEX]
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

func getParamBgColor(ctx *fiber.Ctx) (*img.Color, error) {
	bgColorParam := new(img.Color)
	bgColorParam.R = DEFAULT_BG_COLOR.R
	bgColorParam.G = DEFAULT_BG_COLOR.G
	bgColorParam.B = DEFAULT_BG_COLOR.B

	bgColorValue := ctx.Query(KEY_FOR_BG_COLOR)

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

func getParamTextColor(ctx *fiber.Ctx) (*img.Color, error) {
	txtColorParam := new(img.Color)
	txtColorParam.R = DEFAULT_TEXT_COLOR.R
	txtColorParam.G = DEFAULT_TEXT_COLOR.G
	txtColorParam.B = DEFAULT_TEXT_COLOR.B

	txtColorValue := ctx.Query(KEY_FOR_TEXT_COLOR)

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

func getParamText(ctx *fiber.Ctx, defaultValue string) string {
	return ctx.Query(KEY_FOR_TEXT, defaultValue)
}

func getParamScale(ctx *fiber.Ctx) (float64, error) {
	scaleValue := ctx.Query(KEY_FOR_SCALE)
	if scaleValue == "" {
		return DEFAULT_SCALE, nil
	}
	scaleParam, err := strconv.ParseFloat(scaleValue, 32)
	if err != nil {
		return 0.0, err
	}
	return float64(scaleParam), nil
}
