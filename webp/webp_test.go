package webp

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
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

	buf := bytes.NewBuffer(nil)

	if err := Encode(buf, m, Lossless()); err != nil {
		t.Fatalf("encoding webp: %v", err)
	}

	if err := os.WriteFile(filepath.Join("testdata", "golden-got.webp"), buf.Bytes(), 0o644); err != nil {
		t.Errorf("writing output png: %v", err)
	}

	out, err := Decode(buf)
	if err != nil {
		t.Fatalf("decoding webp: %v", err)
	}

	outb := out.Bounds()
	for xx := outb.Min.X; xx < outb.Max.X; xx++ {
		for yy := outb.Min.Y; yy < outb.Max.Y; yy++ {
			if out.At(xx, yy) != m.At(xx, yy) {
				t.Errorf("color mismatch after lossless encode: (%d, %d)", xx, yy)
			}
		}
	}
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
			t.Errorf("error: %v (x0: %d, y0: %d, x1: %d, y1: %d)", err, x0, y0, x1, y1)
		}
	})
}
