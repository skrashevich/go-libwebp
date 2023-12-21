package webp

import (
	_ "embed"

	"github.com/ebitengine/purego"
)

const libraryName = "libwebp.so"

func dlopen(name string) (uintptr, error) {
	return purego.Dlopen(libraryName, purego.RTLD_NOW|purego.RTLD_GLOBAL)
}
