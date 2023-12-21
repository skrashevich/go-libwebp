package webp

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
)

var (
	_libwebp uintptr

	_WebPEncodeLosslessRGBA func(uintptr, uintptr, uintptr, uintptr, uintptr) uintptr
	_WebPEncodeRGBA         func(uintptr, uintptr, uintptr, uintptr, uintptr, uintptr) uintptr
	_WebPFree               func(uintptr)
)

func loadLibrary() (uintptr, error) {
	handle, err := dlopen(libraryName)
	if err != nil {
		return 0, fmt.Errorf("loading library: %w", err)
	}
	return handle, nil
}

func Init() (err error) {
	if _libwebp != 0 {
		return nil
	}

	_libwebp, err = loadLibrary()
	if err != nil {
		return err
	}

	_WebPEncodeLosslessRGBA = func() (fn func(uintptr, uintptr, uintptr, uintptr, uintptr) uintptr) {
		purego.RegisterLibFunc(&fn, _libwebp, "WebPEncodeLosslessRGBA")
		return
	}()

	_WebPEncodeRGBA = func() (fn func(uintptr, uintptr, uintptr, uintptr, uintptr, uintptr) uintptr) {
		purego.RegisterLibFunc(&fn, _libwebp, "WebPEncodeRGBA")
		return
	}()

	_WebPFree = func() (fn func(uintptr)) {
		purego.RegisterLibFunc(&fn, _libwebp, "WebPFree")
		return
	}()

	return nil
}

func EncodeImpl(
	in unsafe.Pointer,
	w int32,
	h int32,
	bps int32,
	q float32,
	out unsafe.Pointer,
) uint64 {
	Init()
	var sz uintptr
	if q >= 1.0 {
		sz = _WebPEncodeLosslessRGBA(
			uintptr(in),
			uintptr(w),
			uintptr(h),
			uintptr(bps),
			uintptr(out),
		)
	} else {
		sz = _WebPEncodeRGBA(
			uintptr(in),
			uintptr(w),
			uintptr(h),
			uintptr(bps),
			uintptr(q),
			uintptr(out),
		)
	}
	return uint64(sz)
}

func FreeImpl(ptr unsafe.Pointer) {
	_WebPFree(uintptr(ptr))
}
