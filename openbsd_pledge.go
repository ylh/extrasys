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

// Pledge implements its respective syscall. For more information see pledge(2).
func Pledge(promises string, paths []string) (err error) {
	promisesp_, err := syscall.BytePtrFromString(promises)
	if err != nil {
		return
	}
	promisesp, pathsp := unsafe.Pointer(promisesp_), unsafe.Pointer(nil)
	if paths != nil {
		var pathsp_ []*byte
		if pathsp_, err = syscall.SlicePtrFromStrings(paths); err != nil {
			return
		}
		pathsp = unsafe.Pointer(&pathsp_[0])
	}
	_, _, e := syscall.Syscall(SYS_PLEDGE, uintptr(promisesp), uintptr(pathsp), 0)
	use(promisesp)
	use(pathsp)
	return e
}
