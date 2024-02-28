package webp

import (
	_ "embed"

	"golang.org/x/sys/windows"
)

const libraryName = "libwebp.dll"

func dlopen(name string) (uintptr, error) {
	dll, err := windows.LoadDLL(libraryName)
	if err != nil {
		return 0, err
	}
	return uintptr(dll.Handle), nil
}
