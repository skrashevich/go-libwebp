# go-libwebp

This is a [`ccgo`-based](https://pkg.go.dev/modernc.org/ccgo/v3) translation from [libwebp](https://github.com/webmproject/libwebp/)
in the hope to bring webp codec into Go space.

`go run . megopher.png` 

Can consume JPEG and PNG images.

Status: Experimental

## Contributors

This repo represents a joint effort between Chris Waldon (~whereswaldon) and myself.
Thanks Chris for figuring out the build process for compiling libwebp to Go!

We owe a great debt to [@cznic](https://gitlab.com/cznic) for both creating `ccgo` and fixing numerous small
problems we encounted while compiling `libwebp`.

## Todo

- [ ] high level tests for various image formats
- [ ] porcelain wrapper package to integrate with image.Decode
