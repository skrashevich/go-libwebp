//go:build linux && cgo

package webp

import (
	"image"
	"io"

	"git.sr.ht/~jackmordaunt/go-libwebp/lib/common"
)

// HACK: always use lossless when cgo is activated on linux. cgo on linux provokes
// a bug in purego when using float arguments.
func encodeImpl(w io.Writer, m *image.NRGBA, quality float32) error {
	return common.Encode(w, m, quality, wrappedLossless, WebPFree)
}
