package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"strings"
)

type Flag int

func (f *Flag) GetInfo(arg int, result *int) error {
	*result = 1
	return nil
}

func main() {
	fmt.Println("enter vro")
	Flag := new(Flag)
	rpc.Register(Flag)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":12306")
	if err != nil {
		log.Fatal("listen failed, the port may be occupied:", err)
	}
	http.Serve(l, nil)
}
