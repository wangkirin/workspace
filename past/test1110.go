package main

import (
	"os"

	libseccomp "github.com/seccomp/libseccomp-golang"
)

func main() {
	// scmpsyscall, _ := libseccomp.GetSyscallFromName("getwd")
	filter, _ := libseccomp.NewFilter(libseccomp.ActErrno)
	file, _ := os.Create("/root/seccomp/test")
	filter.ExportPFC(file)
	defer file.Close()

}
