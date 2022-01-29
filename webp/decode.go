package webp

import (
	"image"
	"io"

	stdwebp "golang.org/x/image/webp"
)

// Decode reads a WEBP image from r and returns it as an image.Image.
func Decode(r io.Reader) (image.Image, error) {
	return stdwebp.Decode(r)
}

// DecodeConfig returns the color model and dimensions of a WEBP image without
// decoding the entire image.
func DecodeConfig(r io.Reader) (image.Config, error) {
	return stdwebp.DecodeConfig(r)
}
