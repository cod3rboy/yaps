# <div><span>YA</span><img src="./assets/logo.svg" width="40" height="40" alt="logo" align="center"><span>S</span></div>

Yet Another Placeholder Service

YAPS is a server backend which generates placeholder image of any size along with parameters for customization.

## Installation

`go install github.com/cod3rboy/yaps`

## Docker Image

`docker run -d -p 80:8080 --name yaps-server cod3rboy/yaps`

## Project Dependencies

- [go-webpbin](https://pkg.go.dev/github.com/nickalie/go-webpbin) - For webp format encoding.
- [gg](https://github.com/fogleman/gg) - Go graphics library.
- [gofiber](https://github.com/gofiber/fiber/v2) - Web framework for golang.
- [freetype](https://github.com/golang/freetype) - Font rendering library.
- [iniflags](https://github.com/vharitonsky/iniflags) - Library to load flags from ini configuration files.
- [image](https://pkg.go.dev/golang.org/x/image) - Supplementary library to standard `image` package.

## Building Project

`go build .`

## Running Server

`yaps -hostName localhost -hostPort 8080`

## Running Tests

`go test ./... -v`

## Command Line Flags

| Flag           | Description                                     | Default                              |
| -------------- | ----------------------------------------------- | ------------------------------------ |
| `hostName`     | Server host name to use.                        | `localhost`                          |
| `hostPort`     | Server port number to use.                      | `8080`                               |
| `pathPrefix`   | Prefix path for all routes.                     | `/`                                  |
| `allowMethods` | Comma-separated http methods to allow for CORS. | `GET,PUT,PATCH,POST`                 |
| `allowOrigins` | Commad-separated whitelisted origins for CORS.  | `example.com,foo.com,bar.com` or `*` |
| `config`       | Path to ini configuration file.                 |                                      |

## Docker Image Environment Variables

- `PORT`: Listener port for yaps server.
- `PATH_PREFIX` : Path prefix for all routes.

## Usage Guide

Once the server is up and running (at localhost:8080 for example) then you can send `GET` request to generate and receive placeholder images. The request _path_ determines the image format and _query parameters_ are used to customize image -

| Path  | Image Format |
| ----- | ------------ |
| /png  | PNG Image    |
| /jpg  | JPG Image    |
| /jpeg | JPEG Image   |
| /tiff | TIFF Image   |
| /webp | WEBP Image   |

| Query Parameter | Description                            | Example       |
| --------------- | -------------------------------------- | ------------- |
| s               | Image dimensions (width x height)      | 200x100       |
| b               | Background color in hexadecimal digits | F3FFEA or FA3 |
| c               | Text color in hexadecimal digits       | F3FFEA or FA3 |
| t               | Text to display in the image           | Hello World   |
| x               | Scaling factor for width and height    | 2 or 1.5      |

## Examples

Default image (No query parameters)

![Default Image](assets/examples/default.png)

Custom size (`?s=200x150`)

![Custom Size](assets/examples/custom_size.png)

Yellow background (`?s=200x150&b=FFFF00`)

![Yellow Background](assets/examples/yellow_background.png)

Red text (`?s=200x150&b=FFFF00&c=FF0000`)

![Red Text](assets/examples/red_text.png)

Custom text (`?s=200x150&b=FFFF00&c=FF0000&t=Hello%20World`)

![Custom Text](assets/examples/custom_text.png)

Scaled image (`?s=200x100&b=22DB9B&c=FFF&x=1.5`)

![Scaled Image](assets/examples/scaled_image.png)
