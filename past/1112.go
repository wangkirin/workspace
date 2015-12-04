package main

// #cgo LDFLAGS: -ldl
// #include <stdlib.h>
// #include <unistd.h>
// #include <dlfcn.h>
//
// int
// getcon(void *f, char **con, char **mode)
// {
//   int (*aa_getcon)(char **, char **);
//   aa_getcon= (int (*)(char **, char **))f;
//   return aa_getcon(con, mode);
// }
//
import "C"
import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	libname := C.CString("libapparmor.so.1")
	defer C.free(unsafe.Pointer(libname))
	handle := C.dlopen(libname, C.RTLD_LAZY)
	if handle == nil {
		// we can't open libapparmor.so so we assume libapparmor-dev is not
		// installed and we're not running from a unit file
		fmt.Println("libapparmor.so does not exsit")
		return
	}
	defer func() {
		if r := C.dlclose(handle); r != 0 {
			fmt.Errorf("error closing libapparmor.so")
		}
	}()

	getconfunc := C.dlsym(handle, C.CString("aa_getcon"))
	if getconfunc == nil {
		fmt.Errorf("error resolving aa_getcon function")
		return
	}
	cmode := C.CString("")
	ccon := C.CString("")
	defer C.free(unsafe.Pointer(cmode))
	defer C.free(unsafe.Pointer(ccon))
	if _, err := C.getcon(getconfunc, &ccon, &cmode); err != nil {
		fmt.Println(err)
	}
	aacon := C.GoString(ccon)
	aamode := C.GoString(cmode)
	if !strings.EqualFold(aamode, "enforce") {
		fmt.Println("mybe container not support apprmor")
	}
	if strings.EqualFold(aacon, "runc-test") {
		fmt.Println("apparmor-pass")
	}

}
