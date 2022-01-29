package main

import (
	"flag"
	"fmt"
	"image"
	"os"

	_ "image/jpeg"
	"image/png"
	_ "image/png"

	"git.sr.ht/~jackmordaunt/go-libwebp/webp"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func run(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("not enough arguments")
	}
	switch args[0] {
	case "decode":
		return decodeWebp(args[1:])
	case "encode":
		return encodeWebp(args[1:])
	default:
		return fmt.Errorf("unknown command: %q", args[0])
	}
}

func encodeWebp(args []string) error {
	var (
		quality  float64
		lossless bool
		output   string
	)
	cli := flag.NewFlagSet("encode", flag.ExitOnError)
	cli.BoolVar(&lossless, "lossless", false, "lossless quality, ignores quality flag")
	cli.Float64Var(&quality, "q", 1.0, "quality from [0,1]")
	cli.StringVar(&output, "o", "out.webp", "path to output file")
	if err := cli.Parse(args); err != nil {
		return err
	}
	if len(cli.Args()) < 1 {
		return fmt.Errorf("provide image to encode as webp, try megopher.png")
	}
	input := cli.Args()[0]
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

func decodeWebp(args []string) error {
	var (
		output string
	)
	cli := flag.NewFlagSet("decode", flag.ExitOnError)
	cli.StringVar(&output, "o", "out.png", "path to output file")
	if err := cli.Parse(args); err != nil {
		return err
	}
	if len(cli.Args()) < 1 {
		return fmt.Errorf("provide image to encode as webp")
	}
	input := cli.Args()[0]
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
	img, err := webp.Decode(srcf)
	if err != nil {
		return fmt.Errorf("decoding src image: %w", err)
	}
	if err := png.Encode(dstf, img); err != nil {
		return fmt.Errorf("encoding webp: %w", err)
	}
	return nil
}
