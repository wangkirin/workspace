package main

import (
	"fmt"
	"os"
)

func main() {
	devinfo, _ := os.Stat("/root")
	devmode := devinfo.Mode()
	fmt.Println(devmode)
	fmt.Println(os.FileMode.Perm(devmode))
	fmt.Println(os.ModeDevice)
	fmt.Println(os.ModeCharDevice)
	fmt
}
