package main

import (
	"crypto/md5"
	"fmt"
)

func encrypt(source []byte) string {
	result := md5.Sum(source)
	return fmt.Sprintf("%x", result)
}


func func main() {
	
}() {
	
}