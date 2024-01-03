//go:build !(linux && cgo)

package webp

import (
	"image"
	"io"

	"git.sr.ht/~jackmordaunt/go-libwebp/lib/common"
)

func encodeImpl(w io.Writer, m *image.RGBA, quality float32) error {
	if quality >= 1.0 {
		return common.Encode(w, m, quality, wrappedLossless, WebPFree)
	}
	return common.Encode(w, m, quality, WebPEncodeRGBA, WebPFree)
}
