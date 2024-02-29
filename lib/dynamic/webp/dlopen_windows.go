//go:build windows

package webp

import "golang.org/x/sys/windows"

func dlopen(name string) (uintptr, error) {
	dll, err := windows.LoadDLL(libraryName)
	if err != nil {
		return 0, err
	}
	return uintptr(dll.Handle), nil
}
