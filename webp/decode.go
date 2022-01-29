package webp

import (
	"fmt"
	"image"
	"io"
)

// Decode a webp into an image.
func Decode(r io.Reader) (image.Image, error) {
	return nil, fmt.Errorf("unimplemented")
}

func DecodeConfig(r io.Reader) (image.Config, error) {
	return image.Config{}, fmt.Errorf("unimplemented")
}
