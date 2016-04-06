// +build openbsd
// +build 386 amd64 arm

package extrasys

import (
	"errors"
	"syscall"
	"unsafe"
)

const (
	SYS_PLEDGE = 108
)

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
	_, _, e := syscall.Syscall(SYS_PLEDGE, uintptr(promisesp), 0, 0)
	use(promisesp)
	use(pathsp)
	return errors.New("syscall: " + syscall.Errno(e).Error())
}
