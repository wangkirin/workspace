package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	queue := make(chan int, 1)
	go Consumer(queue)
	time.Sleep(1 * time.Second)
	go Producer(queue)
	time.Sleep(3 * time.Second)
}

func Producer(queue chan<- int) {
	fmt.Println("enter Producer")
	/*	for i := 0; i < 10; i++ {
		fmt.Printf("Produce:%v \n", i)
		queue <- i
	}*/
	queue <- 1
}

func Consumer(queue chan<- int) {
	var outbytes bytes.Buffer
	fmt.Printf("address=%p \n", &queue)
	queaddr := fmt.Sprintf("%x", &queue)
	cmd := exec.Command("./12011", queaddr)
	cmd.Stdout = &outbytes
	cmd.Run()
	fmt.Println("out=" + outbytes.String())
}
