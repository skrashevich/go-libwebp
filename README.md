# go-libwebp

This is a [`ccgo`-based](https://pkg.go.dev/modernc.org/ccgo/v3) translation from [libwebp](https://github.com/webmproject/libwebp/)
in the hope to bring webp encoder into Go space.

`go run ./cmd encode ./cmd/megopher.png` 

Can consume JPEG and PNG images.

Status: Experimental. The translated code appears to work but has not been rigously tested.

## Rationale

There exists a webp decoder for Go `golang.org/x/image/webp` that does not include an encoder. Using it is obtuse: if you need an encoder you need to hack together another dependency.
This package extends `golang.org/x/image/webp` with an encoder translated from the libwebp project, providing a single package webp codec.

## Contributors

This repo represents a joint effort between Chris Waldon (~whereswaldon) and myself.
Thanks Chris for figuring out the build process for compiling libwebp to Go!

We owe a great debt to [@cznic](https://gitlab.com/cznic) for both creating `ccgo` and fixing numerous small
problems we encounted while compiling `libwebp`.

## Todo

- [ ] high level tests for various image formats
- [ ] C style error handling
