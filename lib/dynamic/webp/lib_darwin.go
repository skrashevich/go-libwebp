package webp

import (
	_ "embed"

	"github.com/ebitengine/purego"
)

const libraryName = "libwebp.dylib"

func dlopen(name string) (uintptr, error) {
	return purego.Dlopen(libraryName, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
}
