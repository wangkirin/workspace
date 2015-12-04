package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"runtime"
)

type Host struct {
	MntStat   int
	CheckStat int
}

func InitHost() *Host {
	h := new(Host)
	h.MntStat = 0
	h.CheckStat = 0
	return h
}
func (h *Host) GetMntStat(arg int, result *int) error {
	*result = h.MntStat
	return nil
}

func (h *Host) GetCheckStat(arg int, result *int) error {
	*result = h.CheckStat
	return nil
}

func main() {
	var (
		client           *rpc.Client
		err              error
		constartstat     int
		conmntstat       int
		congetresultstat int
	)
	//init container object
	host := InitHost()
	//set up RPC server and listening
	rpc.Register(host)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":3456")
	if err != nil {
		fmt.Println("listen failed")
	}
	fmt.Println("listening 3456")
	go http.Serve(l, nil)
	runtime.Gosched()
	//get connection with Container
	for {
		client, err = rpc.DialHTTP("tcp", "127.0.0.1:1234")
		if err != nil {
			fmt.Println("link rpc server fial:", err)
		} else {
			break
		}
		time.Sleep(3 * time.Second)
	}
	fmt.Println("connection with container ok")
	// waiting for the start status from Container
	for {
		err = client.Call("Container.GetStartStat", 1, &constartstat)
		if err != nil {
			fmt.Println("call remote service failed", err)
		}
		fmt.Println("remote return success", constartstat)
		if hostmntstat != 0 {
			fmt.Println("get Container start stat")
			break
		}
		time.Sleep(2 * time.Second)
	}

	// waiting for the mount status from Container
	for {
		err = client.Call("Container.GetMntStat", 1, &conmntstat)
		if err != nil {
			fmt.Println("call remote service failed", err)
		}
		fmt.Println("remote return success", conmntstat)
		if hostmntstat != 0 {
			fmt.Println("get Container mnt stat")
			break
		}
		time.Sleep(2 * time.Second)
	}

	// waiting for the mount status from Container
	for {
		err = client.Call("Container.GetMntStat", 1, &conmntstat)
		if err != nil {
			fmt.Println("call remote service failed", err)
		}
		fmt.Println("remote return success", conmntstat)
		if hostmntstat != 0 {
			fmt.Println("get Container mnt stat")
			break
		}
		time.Sleep(2 * time.Second)
	}

}
