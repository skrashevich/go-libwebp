package webp

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	"image/png"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"testing"

	stdwebp "golang.org/x/image/webp"
)

//go:embed testdata/golden-in.png
var goldenIn []byte

//go:embed testdata/golden-out.webp
var goldenOut []byte

func TestLossless(t *testing.T) {
	m, err := png.Decode(bytes.NewReader(goldenIn))
	if err != nil {
		t.Fatalf("decoding image: %v", err)
	}

	buf1 := bytes.NewBuffer(nil)
	buf2 := bytes.NewBuffer(nil)

	if err := Encode(io.MultiWriter(buf1, buf2), m, Lossless()); err != nil {
		t.Fatalf("encoding webp: %v", err)
	}

	if err := os.WriteFile(filepath.Join("testdata", "golden-got.webp"), buf1.Bytes(), 0o644); err != nil {
		t.Errorf("writing output png: %v", err)
	}

	t.Run("stdlib webp outupt", func(t *testing.T) {
		assertOutput(t, m, buf1, stdwebp.Decode)
	})

	t.Run("transpiled webp output", func(t *testing.T) {
		assertOutput(t, m, buf2, Decode)
	})
}

func assertOutput(t *testing.T, m image.Image, src io.Reader, decode func(io.Reader) (image.Image, error)) {
	t.Helper()

	out, err := decode(src)
	if err != nil {
		t.Fatalf("decoding webp: %v", err)
	}

	outb := out.Bounds()
	for xx := outb.Min.X; xx < outb.Max.X; xx++ {
		for yy := outb.Min.Y; yy < outb.Max.Y; yy++ {
			if got, want := out.At(xx, yy), m.At(xx, yy); !colorEqual(got, want) {
				t.Fatalf("color mismatch after lossless encode: Point = (%d, %d) Got = %v Want = %v", xx, yy, got, want)
			}
		}
	}
}

func colorEqual(left, right color.Color) bool {
	lr, lg, lb, la := left.RGBA()
	rr, rg, rb, ra := right.RGBA()
	return lr == rr && lg == rg && lb == rb && la == ra
}

func FuzzEncode(f *testing.F) {
	f.Add(uint16(0), uint16(0), uint16(0), uint16(0), int64(0), float32(0))
	f.Add(uint16(0), uint16(0), uint16(1), uint16(1), int64(1), float32(1))
	f.Add(uint16(0), uint16(0), uint16(1), uint16(1), int64(2), float32(0.5))
	f.Add(uint16(0), uint16(0), uint16(100), uint16(100), int64(3), float32(0.75))
	f.Add(uint16(0), uint16(0), uint16(100), uint16(100), int64(3), float32(0.9))
	f.Add(uint16(0), uint16(0), uint16(100), uint16(100), int64(3), float32(0.95))
	f.Add(uint16(0), uint16(0), uint16(100), uint16(100), int64(3), float32(1.0))
	f.Fuzz(func(t *testing.T, x0, y0, x1, y1 uint16, seed int64, quality float32) {
		rng := rand.New(rand.NewSource(seed))
		m := image.NewNRGBA(image.Rect(int(x0), int(y0), int(x1), int(y1)))
		for x := x0; x < x1; x += 1 {
			for y := y0; y < y1; y += 1 {
				colors := rng.Uint32()
				r := uint8(colors)
				g := uint8(colors >> 8)
				b := uint8(colors >> 16)
				a := uint8(colors >> 24)
				m.Set(int(x), int(y), color.RGBA{R: r, G: g, B: b, A: a})
			}
		}
		if quality <= 0 || quality > 1 {
			t.Skip()
			return
		}
		if int(x1)-int(x0) <= 0 || int(y1)-int(y0) <= 0 {
			t.Skip()
			return
		}
		buf := bytes.NewBuffer(nil)
		if err := Encode(buf, m, Quality(quality)); err != nil {
			t.Errorf("encode error: %v (x0: %d, y0: %d, x1: %d, y1: %d)", err, x0, y0, x1, y1)
		}
		if _, err := Decode(buf); err != nil {
			t.Errorf("decode error: %v (x0: %d, y0: %d, x1: %d, y1: %d)", err, x0, y0, x1, y1)
		}
	})
}
