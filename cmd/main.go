package main

import (
	"flag"
	"fmt"
	"image"
	"os"
	"runtime"

	_ "image/jpeg"
	_ "image/png"

	"git.sr.ht/~jackmordaunt/go-libwebp/webp"
)

func main() {
	runtime.LockOSThread()
	if err := run(os.Args[1:]); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func run(args []string) error {
	var (
		quality  float64
		lossless bool
		output   string
	)
	flag.BoolVar(&lossless, "lossless", false, "lossless quality, ignores quality flag")
	flag.Float64Var(&quality, "q", 1.0, "quality from [0,1]")
	flag.StringVar(&output, "o", "out.webp", "path to output file")
	flag.Parse()
	if len(flag.Args()) < 1 {
		return fmt.Errorf("provide image to encode as webp, try megopher.png")
	}
	input := flag.Args()[0]
	srcf, err := os.Open(input)
	if err != nil {
		return fmt.Errorf("opening input file: %w", err)
	}
	defer srcf.Close()
	dstf, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("opening output file: %w", err)
	}
	defer dstf.Close()
	img, _, err := image.Decode(srcf)
	if err != nil {
		return fmt.Errorf("decoding src image: %w", err)
	}
	if err := webp.Encode(dstf, img); err != nil {
		return fmt.Errorf("encoding webp: %w", err)
	}
	return nil
}
