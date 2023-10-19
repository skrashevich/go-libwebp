package webp

import (
	"bytes"
	_ "embed"
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

//go:embed testdata/golden-in.png
var goldenIn []byte

//go:embed testdata/golden-out.webp
var goldenOut []byte

// TestSmoke runs a basic smoke test to see if we can encode and decode a webp.
func TestSmoke(t *testing.T) {
	m, err := png.Decode(bytes.NewReader(goldenIn))
	if err != nil {
		t.Fatalf("decoding image: %v", err)
	}

	buf := bytes.NewBuffer(nil)

	if err := Encode(buf, m); err != nil {
		t.Fatalf("encoding webp: %v", err)
	}

	if !bytes.Equal(goldenOut, buf.Bytes()) {
		t.Errorf("bytes not equal, want %d, got %d", len(goldenOut), buf.Len())
	}

	if err := os.WriteFile(filepath.Join("testdata", "golden-got.webp"), buf.Bytes(), 0o655); err != nil {
		t.Errorf("writing output png: %v", err)
	}

	if _, err := Decode(buf); err != nil {
		t.Fatalf("decoding webp: %v", err)
	}
}
