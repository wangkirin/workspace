package main

import (
	"fmt"

	"gopkg.in/rainycape/dl.v0"
)

func main() {
	fmt.Println("enter main")
	lib, err := dl.Open("libapparmor.so.1", 0)
	if err != nil {
		fmt.Errorf("load libapparmor failed")
	}
	fmt.Println("open libappa")
	defer lib.Close()
	var getcon func(*string, *string) int
	if err := lib.Sym("aa_getcon", &getcon); err != nil {
		fmt.Println("cannot open aa_getcon")
		panic(err)
	}
	mode := ""
	con := ""
	getcon(&mode, &con)
	fmt.Println(mode + "  ;" + con)
}
