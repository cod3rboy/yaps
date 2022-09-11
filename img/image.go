// Package img provides constants, types and functions for automatic image generation.
package img

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/chai2010/webp"
	"github.com/cod3rboy/yaps/utils"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/tiff"
)

// Constants for image file extensions
const (
	IMAGE_JPG  = "jpg"
	IMAGE_JPEG = "jpeg"
	IMAGE_PNG  = "png"
	IMAGE_TIFF = "tiff"
	IMAGE_WEBP = "webp"
)

// Mime types corresponding to the image extensions
var mimeTypes = map[string]string{
	IMAGE_JPG:  "image/" + IMAGE_JPG,
	IMAGE_JPEG: "image/" + IMAGE_JPEG,
	IMAGE_PNG:  "image/" + IMAGE_PNG,
	IMAGE_TIFF: "image/" + IMAGE_TIFF,
	IMAGE_WEBP: "image/" + IMAGE_WEBP,
}

// ImageEncoderFunc represents encoding function for any image type.
//
// It should encode the [image.Image] and write the data to the [io.Writer].
// If any error occurs during encoding process, it should return that error otherwise it should return nil.
type ImageEncoderFunc func(image.Image, io.Writer) error

// Mapping of an image extension to its corresponding [ImageEncoderFunc] function.
var encoders = map[string]ImageEncoderFunc{
	IMAGE_JPG:  EncodeJPEG,
	IMAGE_JPEG: EncodeJPEG,
	IMAGE_PNG:  EncodePNG,
	IMAGE_TIFF: EncodeTIFF,
	IMAGE_WEBP: EncodeWEBP,
}

// Pixel to Point value scale factor (1px = 0.75pt)
const PX_TO_PT = 0.75

// A Size represents the dimensions of an image i.e. image width and image height.
type Size struct {
	Width  int
	Height int
}

// A Color represents the pixel color of an image.
// Each color component is stored as a separate uint8 value.
type Color struct {
	R uint8
	G uint8
	B uint8
}

// An ImageParams stores parameters for image generation.
type ImageParams struct {
	Format          string  // Image extension
	*Size                   // Image Size
	BackgroundColor *Color  // Color to use for background
	TextColor       *Color  // Color to use for text
	Scale           float64 // Value by which to scale Size
	Text            string  // Text to write on the image
}

// An ImageResult stores data of generated image.
type ImageResult struct {
	Bytes    []byte // Encoded bytes of generated image
	Length   int    // Length = len(Bytes)
	MimeType string // Image mime-type in format "image/<extension>" e.g. image/png, for png image.
}

// Generate generates an image with given parameters.
//
// It returns [ImageResult], nil when image is generated successfully.
// If error occurs while generating image, it returns nil, error.
//
// The actual dimensions of image is calculated by scaling width and height with scale factor in [ImageParams].
func Generate(params *ImageParams) (*ImageResult, error) {
	// Scale dimensions by scale factor
	w := utils.ScaleDimension(params.Width, params.Scale)
	h := utils.ScaleDimension(params.Height, params.Scale)

	canvas := gg.NewContext(w, h)

	// Background filling
	FillBackground(canvas, params.BackgroundColor)
	if err := DrawText(canvas, params.Text, params.TextColor); err != nil {
		return nil, err
	}

	// Determine mime type
	mimeType, exists := mimeTypes[params.Format]
	if !exists {
		return nil, fmt.Errorf("mime type not found for format %s", params.Format)
	}

	// Encode image
	imgBytes, err := encode(canvas, params.Format)
	if err != nil {
		return nil, err
	}

	return &ImageResult{
		Bytes:    imgBytes,
		Length:   len(imgBytes),
		MimeType: mimeType,
	}, nil
}

// FillBackground fills the canvas with given color.
func FillBackground(canvas *gg.Context, color *Color) {
	canvas.SetRGBA255(int(color.R), int(color.G), int(color.B), 0xFF)
	canvas.Clear()
}

// DrawText draws the given text on the canvas with given color.
//
// Dynamic font size is used to draw text and is calcuated by canvas height * [PX_TO_PT] * 0.2.
//
// The text is anchored at the image centre.
// It also wraps around when overflows the canvas width.
func DrawText(canvas *gg.Context, text string, color *Color) error {
	canvas.SetRGBA255(int(color.R), int(color.G), int(color.B), 0xFF)
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return err
	}
	fontFace := truetype.NewFace(font, &truetype.Options{Size: float64(canvas.Height()) * PX_TO_PT * 0.2})
	canvas.SetFontFace(fontFace)
	canvas.DrawStringWrapped(text, float64(canvas.Width()/2), float64(canvas.Height())/2, 0.5, 0.5, float64(canvas.Width())*0.8, 1, gg.AlignCenter)
	return nil
}

// encode converts the canvas into bytes for given image format.
//
// If an error occurs while encoding, it returns nil, error.
func encode(canvas *gg.Context, format string) ([]byte, error) {
	imgBuffer := new(bytes.Buffer)

	encodeFunc, exists := encoders[format]
	if !exists {
		return nil, fmt.Errorf("encoder not found for format %s", format)
	}
	err := encodeFunc(canvas.Image(), imgBuffer)
	if err != nil {
		return nil, err
	}
	return imgBuffer.Bytes(), nil
}

// EncodePNG encodes the given [image.Image] into PNG format bytes and writes those bytes in [io.Writer].
// If an error occurs while encoding or writing, it returns that error.
//
// EncodePNG is compatible with [ImageEncoderFunc] type.
func EncodePNG(img image.Image, writer io.Writer) error {
	return png.Encode(writer, img)
}

// EncodeJPEG encodes the given [image.Image] into JPEG format bytes and writes those bytes in [io.Writer].
// If an error occurs while encoding or writing, it returns that error.
//
// It uses 75% quality while encoding JPEG image.
//
// EncodeJPEG is compatible with [ImageEncoderFunc] type.
func EncodeJPEG(img image.Image, writer io.Writer) error {
	return jpeg.Encode(writer, img, &jpeg.Options{
		Quality: jpeg.DefaultQuality,
	})
}

// EncodeTIFF encodes the given [image.Image] into TIFF format bytes and writes those bytes in [io.Writer].
// If an error occurs while encoding or writing, it returns that error.
//
// It uses deflate compression and a predictor while encoding TIFF image.
//
// EncodeTIFF is compatible with [ImageEncoderFunc] type.
func EncodeTIFF(img image.Image, writer io.Writer) error {
	return tiff.Encode(writer, img, &tiff.Options{
		Compression: tiff.Deflate,
		Predictor:   true,
	})
}

// EncodeWEBP encodes the given [image.Image] into WEBP format bytes and writes those bytes in [io.Writer].
// If an error occurs while encoding or writing, it returns that error.
//
// It uses lossless compression, 90% quality and also preserves transparency while encoding WEBP image.
//
// EncodeWEBP is compatible with [ImageEncoderFunc] type.
func EncodeWEBP(img image.Image, writer io.Writer) error {
	return webp.Encode(writer, img, &webp.Options{
		Lossless: false,
		Quality:  webp.DefaulQuality,
		Exact:    true,
	})
}
