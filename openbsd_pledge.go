// +build openbsd
// +build 386 amd64 arm

package extrasys

import (
	"syscall"
	"unsafe"
)

const (
	SYS_PLEDGE = 108
)

func Pledge(promises string, paths []string) (err error) {
	promisesp, err := syscall.BytePtrFromString(promises)
	if err != nil {
		return
	}
	pathsp, err := syscall.SlicePtrFromStrings(paths)
	if err != nil {
		return
	}
	_, _, e := syscall.Syscall(SYS_PLEDGE, uintptr(unsafe.Pointer(promisesp)), uintptr(unsafe.Pointer(&pathsp[0])), 0)
	use(unsafe.Pointer(promisesp))
	use(unsafe.Pointer(&pathsp[0]))
	return syscall.Errno(e)
}
