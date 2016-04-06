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
	promisesPtr, err := syscall.BytePtrFromString(promises)
	if err != nil {
		return
	}
	promisesUnsafe, pathsUnsafe := unsafe.Pointer(promisesPtr), unsafe.Pointer(nil)
	if paths != nil {
		var pathsPtr []*byte
		if pathsPtr, err = syscall.SlicePtrFromStrings(paths); err != nil {
			return
		}
		pathsUnsafe = unsafe.Pointer(&pathsPtr[0])
	}
	_, _, e := syscall.Syscall(SYS_PLEDGE, uintptr(promisesUnsafe), uintptr(pathsUnsafe), 0)
	use(promisesUnsafe)
	use(pathsUnsafe)
	return e
}
