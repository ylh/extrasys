package extrasys // import "ylh.io/extrasys"

import "unsafe"

// Everything below this line is ripped from the standard library source
// <https://golang.org/src/syscall/syscall.go>

// use is a no-op, but the compiler cannot see that it is.
// Calling use(p) ensures that p is kept live until that point.
// This was needed until Go 1.6 to call syscall.Syscall correctly.
// As of Go 1.6 the compiler handles that case automatically.
// The uses and definition of use can be removed early in the Go 1.7 cycle.
//go:noescape
func use(p unsafe.Pointer)
