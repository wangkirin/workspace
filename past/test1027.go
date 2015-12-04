package main

import (
	"fmt"
	"sync"
)

func main() {
	people := []string{"Anna", "Bob", "Cody", "Dave", "Eva"}
	match := make(chan string, 5) // ÎªÒ»¸öÎ´Æ¥ÅäµÄ·¢ËÍ²Ù×÷Ìá¹©¿Õ¼ä
	wg := new(sync.WaitGroup)
	wg.Add(len(people))
	for _, name := range people {
		go Seek(name, match, wg)
	}
	wg.Wait()
	select {
	case name := <-match:
		fmt.Printf("No one received %s¡¯s message.\n", name)
	default:
		// Ã»ÓÐ´ý´¦ÀíµÄ·¢ËÍ²Ù×÷
	}
}

// º¯ÊýSeek ·¢ËÍÒ»¸önameµ½match¹ÜµÀ»ò´Ómatch¹ÜµÀ½ÓÊÕÒ»¸öpeer£¬½áÊøÊ±Í¨Öªwait group
func Seek(name string, match chan string, wg *sync.WaitGroup) {
	select {
	case peer := <-match:
		fmt.Printf("%s sent a message to %s.\n", peer, name)
	case match <- name:
		// µÈ´ýÄ³¸ögoroutine½ÓÊÕÎÒµÄÏûÏ¢
	}
	wg.Done()
}
