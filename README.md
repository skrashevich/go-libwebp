# go-libwebp

Status: Experimental.

This package provides a means to encode webp images.
It does _not_ wrap the entire libwebp API surface.
It _does_ provide a way to avoid CGO.

This package provides several options for interaction with libwebp:
1. transpiled C to Go (ccgo)
2. dynamic libraries  (purego)

Each option is called a "backend".

Package dynamic provides a dynamic binding to libwebp. This requires the consumer to acquire the shared objects.
Package transpiled provides a Go translation of the libwebp source. This typically executes slower.

If the shared objects are available at runtime, dynamic bindings will be preferred.

Package transpiled is a [`ccgo`-based](https://pkg.go.dev/modernc.org/ccgo/v3) translation from [libwebp](https://github.com/webmproject/libwebp/)
in the hope to bring webp encoder into Go space.

## Example

`go run ./cmd encode ./cmd/megopher.png` 

Can consume JPEG and PNG images.

## Rationale

There exists a webp decoder for Go `golang.org/x/image/webp` that does not include an encoder. Using it is obtuse: if you need an encoder you need to hack together another dependency.
This package extends `golang.org/x/image/webp` with an encoder translated from the libwebp project, providing a single package webp codec.

## Toolchain

The transpilation process relies on a cross compile toolchain.

This is not strictly required at the moment, since the transpiled code is
identical between operating systems and architectures (amd64, arm64).

`tools/transpile.sh` expects these compilers on an amd64 Linux box:

- gcc                   (gcc native)
- aarch64-linux-gnu     (gcc)
- x86_64-w64-mingw32    (mingw)
- x86_64-apple-darwin15 (osxcross)

## Dynamic

To make use of the dynamic bindings, ensure you supply the appropriate shared objects at runtime. 
It is the consumer's responsibility to acquire the shared object. This projects contains a script
that will build amd64 shared-objects and place them in `lib/dynamic/webp/blobs` (`tools/compile-dynamic-libraries.sh`).

## Testing

Go side testing involves a lossless compression test; byte-wise comparison against a golden test image.
In addition, there is a fuzz harness that executes against the transpiled code - encoding and subsequently
decoding images of random size and color.

## Features

- support for amd64 Windows, Linux, Darwin
- support for dynamic linking and transpiled backends
  - if you want speed and must avoid CGO, use dynamic linking
  - if you want to avoid CGO without further hassle, use transpilation

## Contributors

This repo represents a joint effort between Chris Waldon (~whereswaldon) and myself.
Thanks Chris for figuring out the build process for compiling libwebp to Go!

We owe a great debt to [@cznic](https://gitlab.com/cznic) for both creating `ccgo` and fixing numerous small
problems we encounted while compiling `libwebp`.

