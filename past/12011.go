package main

import (
	"fmt"
	"os"
	"strconv"
	"unsafe"
)

func main() {
	addr := os.Args[1]
	fmt.Println("enter Consumer2")
	fmt.Printf("addr=%v \n", addr)
	addrint, _ := strconv.ParseUint(addr, 16, 64)
	fmt.Printf("addrint=%x \n", addrint)
	ptr := uintptr(addrint)
	pchan := (*chan int)(unsafe.Pointer(ptr))
	fmt.Printf("pointer=%p \n", pchan)
	_, ok := <-*pchan
	if ok {
		fmt.Println("13123")
	} else {
		fmt.Println("channel closed")
	}
	// for i := range *pchan {
	// 	fmt.Println("print queue")
	// 	fmt.Println(i)
	// }
}
