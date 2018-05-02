# `/!\ ACHTUNG /!\` DEPRECATED `/!\ ACHTUNG /!\`
In what remains the most significant open-source contribution of my short life thus far, the singular call provided within this library [https://github.com/golang/sys/commit/8fd966b47dbdd4faa03de0d06e3d733baeb9a1a9](has been merged into `golang.org/x/sys/unix`). Please use that instead.

# extrasys
github wouldn't shut up until I added this stupid document

This may eventually serve as a general repository for syscalls not provided by Go's syscall package.
Currently it just provides OpenBSD's pledge(2).

To use, do:
````
go get ylh.io/extrasys
````
