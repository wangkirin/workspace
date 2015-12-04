package main

// #cgo LDFLAGS: -lapparmor
// #include <sys/apparmor.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("enter test apparmor")
	cmode := C.CString("")
	ccon := C.CString("")
	cmntpoint := C.CString("")
	defer C.free(unsafe.Pointer(cmode))
	defer C.free(unsafe.Pointer(ccon))
	fmt.Println("finish init ")
	if _, err := C.aa_getcon(&ccon, &cmode); err != nil {
		fmt.Println(err)
	}
	if _, err := C.aa_find_mountpoint(&cmntpoint); err != nil {
		fmt.Println(err)
	}
	fmt.Println(C.GoString(ccon))
	fmt.Println(C.GoString(cmode))
	fmt.Println(C.GoString(cmntpoint))
}
