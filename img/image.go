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

// Image file extensions
const IMAGE_JPG = "jpg"
const IMAGE_JPEG = "jpeg"
const IMAGE_PNG = "png"
const IMAGE_TIFF = "tiff"
const IMAGE_WEBP = "webp"

// const IMAGE_GIF = "gif"
// const IMAGE_SVG = "svg"

// Mime types corresponding to above images
var mimeTypes = map[string]string{
	IMAGE_JPG:  "image/" + IMAGE_JPG,
	IMAGE_JPEG: "image/" + IMAGE_JPEG,
	IMAGE_PNG:  "image/" + IMAGE_PNG,
	IMAGE_TIFF: "image/" + IMAGE_TIFF,
	IMAGE_WEBP: "image/" + IMAGE_WEBP,
	// IMAGE_GIF:  "image/" + IMAGE_GIF,
	// IMAGE_SVG:  "image/" + IMAGE_SVG,
}

// Encoding function type
type ImageEncoderFunc func(image.Image, io.Writer) error

// Format indexed map of encoding functions
var encoders = map[string]ImageEncoderFunc{
	IMAGE_JPG:  encodeJPEG,
	IMAGE_JPEG: encodeJPEG,
	IMAGE_PNG:  encodePNG,
	IMAGE_TIFF: encodeTIFF,
	IMAGE_WEBP: encodeWEBP,
}

// Pixel to Point value scale factor (1px = 0.75pt)
const PX_TO_PT = 0.75

type Size struct {
	Width  int
	Height int
}

type Color struct {
	R uint8
	G uint8
	B uint8
}

type ImageParams struct {
	Format string
	*Size
	BackgroundColor *Color
	TextColor       *Color
	Scale           float64
	Text            string
}

type ImageResult struct {
	Bytes    []byte
	Length   int
	MimeType string
}

func Generate(params *ImageParams) (*ImageResult, error) {
	w := utils.ScaleDimension(params.Width, params.Scale)
	h := utils.ScaleDimension(params.Height, params.Scale)

	canvas := gg.NewContext(w, h)

	drawBackground(canvas, params.BackgroundColor)
	if err := drawText(canvas, params.Text, params.TextColor); err != nil {
		return nil, err
	}

	mimeType, exists := mimeTypes[params.Format]
	if !exists {
		return nil, fmt.Errorf("mime type not found for format %s", params.Format)
	}

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

func drawBackground(canvas *gg.Context, color *Color) {
	canvas.SetRGBA255(int(color.R), int(color.G), int(color.B), 0xFF)
	canvas.Clear()
}

func drawText(canvas *gg.Context, text string, color *Color) error {
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

func encodePNG(img image.Image, writer io.Writer) error {
	return png.Encode(writer, img)
}

func encodeJPEG(img image.Image, writer io.Writer) error {
	return jpeg.Encode(writer, img, &jpeg.Options{
		Quality: jpeg.DefaultQuality,
	})
}

func encodeTIFF(img image.Image, writer io.Writer) error {
	return tiff.Encode(writer, img, &tiff.Options{
		Compression: tiff.Deflate,
		Predictor:   true,
	})
}

func encodeWEBP(img image.Image, writer io.Writer) error {
	return webp.Encode(writer, img, &webp.Options{
		Lossless: false,
		Quality:  webp.DefaulQuality,
		Exact:    true,
	})
}
