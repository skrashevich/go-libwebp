//go:build !windows

package webp

import "github.com/ebitengine/purego"

func dlopen(name string) (uintptr, error) {
	return purego.Dlopen(libraryName, purego.RTLD_LAZY|purego.RTLD_LOCAL)
}
