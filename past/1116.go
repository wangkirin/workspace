package main

import (
	"fmt"
	"syscall"
)

func main() {
	out, err := syscall.Readlink("/proc/self/ns/", buf)
}
