package img

import (
	"bytes"
	"image"
	"os"
	"path/filepath"
	"testing"

	"github.com/fogleman/gg"
)

func TestFillBackground(t *testing.T) {
	x1, y1, x2, y2 := 0, 0, 100, 100
	ctx := gg.NewContextForRGBA(image.NewRGBA(image.Rect(x1, y1, x2, y2)))
	bgColor := &Color{
		R: 0x35,
		G: 0x6E,
		B: 0xF3,
	}
	FillBackground(ctx, bgColor)
	img := ctx.Image()
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			r &= 0x000000FF
			g &= 0x000000FF
			b &= 0x000000FF
			if r != uint32(bgColor.R) || g != uint32(bgColor.G) || b != uint32(bgColor.B) {
				t.Fatalf("at point (%d, %d) - expected=(%d,%d,%d), actual=(%d,%d,%d)", x, y, bgColor.R, bgColor.G, bgColor.B, r, g, b)
			}
		}
	}
}

type TestData struct {
	Params       ImageParams
	ExpectedPath string
}

var testData = []TestData{
	{
		Params: ImageParams{
			Format:          "jpg",
			Size:            &Size{100, 100},
			BackgroundColor: &Color{0x2D, 0x64, 0xDD},
			TextColor:       &Color{0, 0, 0},
			Scale:           1,
			Text:            "",
		},
		ExpectedPath: "testdata/1.jpg",
	},
	{
		Params: ImageParams{
			Format:          "tiff",
			Size:            &Size{100, 100},
			BackgroundColor: &Color{0, 0, 0},
			TextColor:       &Color{0xFF, 0xFF, 0xFF},
			Scale:           1,
			Text:            "Hello",
		},
		ExpectedPath: "testdata/2.tif",
	},
	{
		Params: ImageParams{
			Format:          "png",
			Size:            &Size{100, 100},
			BackgroundColor: &Color{0xFF, 0xFF, 0xFF},
			TextColor:       &Color{0, 0, 0},
			Scale:           1,
			Text:            "World",
		},
		ExpectedPath: "testdata/3.png",
	},
	{
		Params: ImageParams{
			Format:          "jpeg",
			Size:            &Size{50, 50},
			BackgroundColor: &Color{0xFF, 0xFF, 0xFF},
			TextColor:       &Color{0, 0, 0},
			Scale:           2,
			Text:            "",
		},
		ExpectedPath: "testdata/4.jpeg",
	},
	{
		Params: ImageParams{
			Format:          "webp",
			Size:            &Size{50, 50},
			BackgroundColor: &Color{0xFF, 0xFF, 0xFF},
			TextColor:       &Color{0, 0, 0},
			Scale:           2,
			Text:            "Hello",
		},
		ExpectedPath: "testdata/5.webp",
	},
}

func TestGenerate(t *testing.T) {
	wdir, _ := os.Getwd()
	for _, data := range testData {
		result, err := Generate(&data.Params)
		if err != nil {
			t.Fatal(err)
		}
		expectedMime, actualMime := mimeTypes[data.Params.Format], result.MimeType
		if expectedMime != actualMime {
			t.Fatalf("expected mime type = %s , actual mime type = %s", expectedMime, actualMime)
		}

		expectedPath := filepath.Join(wdir, data.ExpectedPath)
		actualPath := filepath.Join(wdir, "testgen", filepath.Base(expectedPath))

		expectedBytes, err := os.ReadFile(expectedPath)
		if err != nil {
			t.Fatal(err)
		}
		actualBytes := result.Bytes

		os.RemoveAll(filepath.Dir(actualPath))
		err = os.MkdirAll(filepath.Dir(actualPath), 0644)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(expectedBytes, actualBytes) {
			err := os.WriteFile(actualPath, actualBytes, 0644)
			if err != nil {
				t.Fatal(err)
			}
			t.Fatalf("\nexpected image = %s\nactual image = %s\n", expectedPath, actualPath)
		}
	}
}
