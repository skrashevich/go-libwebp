package webp

import (
	"fmt"
	"image"
	"io"
	"runtime"
	"unsafe"

	dynamic "git.sr.ht/~jackmordaunt/go-libwebp/lib/dynamic/webp"
	transpiled "git.sr.ht/~jackmordaunt/go-libwebp/lib/transpiled/webp"
)

// Encode an image into webp with default settings.
func Encode(w io.Writer, m image.Image, opt ...EncodeOption) error {
	var enc Encoder
	for _, op := range opt {
		op(&enc)
	}
	if enc.Quality <= 0.0 {
		enc.Quality = 0.1
	}
	if enc.Quality > 1.0 {
		enc.Quality = 1.0
	}
	if enc.Lossless {
		enc.Quality = 1.0
	}
	return enc.Encode(w, m)
}

// EncodeOption configures the encoder.
type EncodeOption func(*Encoder)

// Quality in the range (0,1].
// Quality of 1 implies Lossless.
func Quality(q float32) EncodeOption {
	return func(enc *Encoder) {
		enc.Quality = q
	}
}

// Lossless will ignore quality.
func Lossless() EncodeOption {
	return func(enc *Encoder) {
		enc.Lossless = true
	}
}

// Encoder implements webp encoding of an image.
type Encoder struct {
	// Quality is in the range (0,1]. Values outside of this
	// range will be treated as 1.
	Quality float32
	// Lossless indicates whether to use the lossless compression
	// strategy. If true, the Quality field is ignored.
	Lossless bool
}

// Encode specified image as webp to w.
// If the image is RGBA, the pixel buffer will be encoded directly.
// If the image is not RGBA, it will be converted to RGBA first.
// Not threadsafe.
func (enc *Encoder) Encode(w io.Writer, m image.Image) error {
	if enc.Quality <= 0.0 || enc.Quality > 1 {
		enc.Quality = 1.0
	}
	if rgba, ok := m.(*image.RGBA); ok {
		return enc.encode(w, rgba)
	}
	rgba := image.NewRGBA(m.Bounds())
	b := m.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			rgba.Set(x, y, m.At(x, y))
		}
	}
	return enc.encode(w, rgba)
}

func (enc *Encoder) encode(w io.Writer, m *image.RGBA) error {
	if err := dynamic.Init(); err == nil {
		return encode(w, m, enc.Quality, dynamic.EncodeImpl, dynamic.FreeImpl)
	}
	return encode(w, m, enc.Quality, transpiled.EncodeImpl, transpiled.FreeImpl)
}

// encode an RGBA image losslessly into the provided writer.
func encode(w io.Writer, m *image.RGBA, q float32, enc encodeFunc, free freeFunc) error {
	out := new(*byte)

	p := runtime.Pinner{}
	defer p.Unpin()

	p.Pin(&m.Pix[0])
	p.Pin(out)

	size := enc(
		unsafe.Pointer(&m.Pix[0]),
		int32(m.Bounds().Dx()),
		int32(m.Bounds().Dy()),
		int32(m.Stride),
		q,
		unsafe.Pointer(out),
	)

	if size == 0 {
		return fmt.Errorf("empty result")
	}

	if *out == nil {
		return fmt.Errorf("failed to allocate output buffer")
	}

	defer free(unsafe.Pointer(*out))

	buf := unsafe.Slice(*out, size)

	if _, err := w.Write(buf); err != nil {
		return fmt.Errorf("writing the encoded webp: %w", err)
	}

	return nil
}
