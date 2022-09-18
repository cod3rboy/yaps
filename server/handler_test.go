package server

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
)

type TestData struct {
	Route              string
	Query              string
	ExpectedStatusCode int
	ExpectedImagePath  string
	ExpectedType       string
}

var testData = []TestData{
	{
		Route:              "/",
		Query:              "",
		ExpectedStatusCode: 404,
		ExpectedImagePath:  "",
		ExpectedType:       "",
	},
	{
		Route:              "/abcd",
		Query:              "",
		ExpectedStatusCode: 404,
		ExpectedImagePath:  "",
		ExpectedType:       "",
	},
	{
		Route:              "/png",
		Query:              "",
		ExpectedStatusCode: 200,
		ExpectedImagePath:  "testdata/1.png",
		ExpectedType:       "image/png",
	},
	{
		Route:              "/jpg",
		Query:              "",
		ExpectedStatusCode: 200,
		ExpectedImagePath:  "testdata/2.jpg",
		ExpectedType:       "image/jpg",
	},
	{
		Route:              "/jpeg",
		Query:              "",
		ExpectedStatusCode: 200,
		ExpectedImagePath:  "testdata/3.jpeg",
		ExpectedType:       "image/jpeg",
	},
	{
		Route:              "/tiff",
		Query:              "",
		ExpectedStatusCode: 200,
		ExpectedImagePath:  "testdata/4.tif",
		ExpectedType:       "image/tiff",
	},
	{
		Route:              "/webp",
		Query:              "",
		ExpectedStatusCode: 200,
		ExpectedImagePath:  "testdata/5.webp",
		ExpectedType:       "image/webp",
	},
	{
		Route:              "/png",
		Query:              "s=110x80&b=5EF3AB&c=FF236D&t=Hello World&x=2",
		ExpectedStatusCode: 200,
		ExpectedImagePath:  "testdata/6.png",
		ExpectedType:       "image/png",
	},
	{
		Route:              "/jpg",
		Query:              "s=110x80&b=5EF3AB&c=FF236D&t=Hello World&x=2",
		ExpectedStatusCode: 200,
		ExpectedImagePath:  "testdata/7.jpg",
		ExpectedType:       "image/jpg",
	},
	{
		Route:              "/jpeg",
		Query:              "s=110x80&b=5EF3AB&c=FF236D&t=Hello World&x=2",
		ExpectedStatusCode: 200,
		ExpectedImagePath:  "testdata/8.jpeg",
		ExpectedType:       "image/jpeg",
	},
	{
		Route:              "/tiff",
		Query:              "s=110x80&b=5EF3AB&c=FF236D&t=Hello World&x=2",
		ExpectedStatusCode: 200,
		ExpectedImagePath:  "testdata/9.tif",
		ExpectedType:       "image/tiff",
	},
	{
		Route:              "/webp",
		Query:              "s=110x80&b=5EF3AB&c=FF236D&t=Hello World&x=2",
		ExpectedStatusCode: 200,
		ExpectedImagePath:  "testdata/10.webp",
		ExpectedType:       "image/webp",
	},
	{
		Route:              "/png",
		Query:              "b=F3D&c=33F&t=Hello World",
		ExpectedStatusCode: 200,
		ExpectedImagePath:  "testdata/11.png",
		ExpectedType:       "image/png",
	},
	{
		Route:              "/png",
		Query:              "x=1.5",
		ExpectedStatusCode: 200,
		ExpectedImagePath:  "testdata/12.png",
		ExpectedType:       "image/png",
	},
	{
		Route:              "/png",
		Query:              "x=.5",
		ExpectedStatusCode: 200,
		ExpectedImagePath:  "testdata/13.png",
		ExpectedType:       "image/png",
	},
	{
		Route:              "/png",
		Query:              "s=100+23",
		ExpectedStatusCode: 500,
		ExpectedImagePath:  "",
		ExpectedType:       "",
	},
	{
		Route:              "/png",
		Query:              "b=F3FF",
		ExpectedStatusCode: 500,
		ExpectedImagePath:  "",
		ExpectedType:       "",
	},
	{
		Route:              "/png",
		Query:              "b=FA35AZ",
		ExpectedStatusCode: 500,
		ExpectedImagePath:  "",
		ExpectedType:       "",
	},
	{
		Route:              "/png",
		Query:              "c=F3FF",
		ExpectedStatusCode: 500,
		ExpectedImagePath:  "",
		ExpectedType:       "",
	},
	{
		Route:              "/png",
		Query:              "c=FA35AZ",
		ExpectedStatusCode: 500,
		ExpectedImagePath:  "",
		ExpectedType:       "",
	},
	{
		Route:              "/png",
		Query:              "x=i",
		ExpectedStatusCode: 500,
		ExpectedImagePath:  "",
		ExpectedType:       "",
	},
}

func TestHandlerImage(t *testing.T) {
	router := fiber.New()
	router.Get("/:format<regex("+strings.Join(SupportedFormats, "|")+")>", HandlerImage)

	for _, data := range testData {
		reqUrl := data.Route
		if data.Query != "" {
			reqUrl += "?" + url.PathEscape(data.Query)
		}
		req := httptest.NewRequest(http.MethodGet, reqUrl, nil)
		res, err := router.Test(req)
		if err != nil {
			t.Fatal(err)
		}
		actualStatusCode := res.StatusCode
		if actualStatusCode != res.StatusCode {
			t.Fatalf("\nexpected status code = %d\nactual status code = %d\n", data.ExpectedStatusCode, actualStatusCode)
		}
		if data.ExpectedImagePath != "" {
			expectedType, actualType := data.ExpectedType, res.Header.Get("content-type")
			if expectedType != actualType {
				t.Fatalf("\nexpected content type = %s\nactual content type = %s\n", expectedType, actualType)
			}
			wdir, _ := os.Getwd()
			expectedPath := filepath.Join(wdir, data.ExpectedImagePath)
			actualPath := filepath.Join(wdir, "testgen", filepath.Base(expectedPath))

			expectedBytes, err := os.ReadFile(expectedPath)
			if err != nil {
				t.Fatal(err)
			}
			actualBytes, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

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
}
