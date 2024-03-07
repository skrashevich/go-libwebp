package webp

import (
	"fmt"
	"image"
	"io"
	"runtime"

	"git.sr.ht/~jackmordaunt/go-libwebp/lib/common"
	"github.com/ebitengine/purego"
)

var (
	_libwebp uintptr

	WebPEncodeLosslessRGBA func(in uintptr, width int32, height int32, bps int32, out uintptr) (size uint64)
	WebPEncodeRGBA         func(in uintptr, width int32, height int32, bps int32, quality float32, out uintptr) (size uint64)
	WebPDecodeRGBA         func(data uintptr, size uint64, width uintptr, height uintptr) uintptr
	WebPFree               func(ptr uintptr)
)

func genericLibraryName() string {
	switch runtime.GOOS {
	case "windows":
		return "libwebp.dll"
	case "darwin":
		return "libwebp.dylib"
	default:
		return "libwebp.so"
	}
}

func loadLibrary() (uintptr, error) {
	handle, err := dlopen(libraryName)
	if err != nil {
		handle, err = dlopen(genericLibraryName())
	}
	return handle, err
}

func Init() (err error) {
	if _libwebp != 0 {
		return nil
	}

	_libwebp, err = loadLibrary()
	if err != nil {
		return fmt.Errorf("loading library: %w", err)
	}

	purego.RegisterLibFunc(&WebPEncodeLosslessRGBA, _libwebp, "WebPEncodeLosslessRGBA")
	purego.RegisterLibFunc(&WebPEncodeRGBA, _libwebp, "WebPEncodeRGBA")
	purego.RegisterLibFunc(&WebPDecodeRGBA, _libwebp, "WebPDecodeRGBA")
	purego.RegisterLibFunc(&WebPFree, _libwebp, "WebPFree")

	return nil
}

func DecodeImpl(buf []byte) (image.Image, error) {
	return common.Decode(buf, WebPDecodeRGBA, WebPFree)
}

func EncodeImpl(w io.Writer, m *image.NRGBA, quality float32) error {
	return encodeImpl(w, m, quality)
}

// wrappedLossless drops the quality param and does a lossless encode.
func wrappedLossless(in uintptr, w int32, h int32, bps int32, _ float32, out uintptr) uint64 {
	return WebPEncodeLosslessRGBA(in, w, h, bps, out)
}
