package webp

import (
	"unsafe"

	"modernc.org/libc"
)

var tls *libc.TLS

func EncodeImpl(
	in unsafe.Pointer,
	w int32,
	h int32,
	bps int32,
	q float32,
	out unsafe.Pointer,
) (size uint64) {
	tls = libc.NewTLS()
	defer tls.Close()

	if q >= 1.0 {
		size = WebPEncodeLosslessRGBA(
			tls,
			uintptr(in),
			w,
			h,
			bps,
			uintptr(out),
		)
	} else {
		size = WebPEncodeRGBA(
			tls,
			uintptr(in),
			w,
			h,
			bps,
			q,
			uintptr(out),
		)
	}

	return size
}

func FreeImpl(ptr unsafe.Pointer) {
	if tls == nil {
		panic("tls is nil, invalid use of transpile api")
	}
	WebPFree(tls, uintptr(ptr))
}
