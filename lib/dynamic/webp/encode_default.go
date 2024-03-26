//go:build !(linux && cgo)

package webp

import (
	"fmt"
	"image"
	"io"

	"git.sr.ht/~jackmordaunt/go-libwebp/lib/common"
)

func encodeImpl(w io.Writer, m *image.NRGBA, quality float32) error {
	if quality >= 1.0 {
		return common.Encode(w, m, quality, wrappedLossless, WebPFree)
	}
	if WebPEncodeRGBA == nil || WebPFree == nil {
		return fmt.Errorf("functions not initialized")
	}
	return common.Encode(w, m, quality, WebPEncodeRGBA, WebPFree)
}
