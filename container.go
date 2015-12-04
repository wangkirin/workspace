package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"runtime"
	"time"
)

type Container struct {
	StartStat   int
	MntStat     int
	IsResultGet int
}

func InitContainer() *Container {
	con := new(Container)
	con.StartStat = 0
	con.MntStat = 0
	con.IsResultGet = 0
	return con
}

func (c *Container) GetStartStat(arg int, result *int) error {
	*result = c.StartStat
	return nil
}

func (c *Container) GetMntStat(arg int, result *int) error {
	*result = c.MntStat
	return nil
}

func (c *Container) GetIsResultGet(arg int, result *int) error {
	*result = c.IsResultGet
	return nil
}

func main() {
	var (
		client        *rpc.Client
		err           error
		hostmntstat   int
		hostcheckstat int
	)
	//init container object
	container := InitContainer()
	//container started ,set StartStat non-zero
	container.StartStat = 1
	//set up RPC server and listening
	rpc.Register(container)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("listen failed")
	}
	fmt.Println("listening 1234")
	go http.Serve(l, nil)
	runtime.Gosched()
	//get connection with Host
	for {
		client, err = rpc.DialHTTP("tcp", "127.0.0.1:3456")
		if err != nil {
			fmt.Println("link rpc server fial:", err)
		} else {
			break
		}
		time.Sleep(3 * time.Second)
	}
	fmt.Println("connection  out ok")
	// waiting for the mount information message from Host
	for {
		err = client.Call("Host.GetMntStat", 1, &hostmntstat)
		if err != nil {
			fmt.Println("call remote service failed", err)
		}
		fmt.Println("remote return success", hostmntstat)
		if hostmntstat != 0 {
			fmt.Println("get mnt stat")
			break
		}
		time.Sleep(3 * time.Second)
	}
	// TODO:mnt in the contianer
	fmt.Println("mnt inside container finish")
	container.MntStat = 1
	// waiting for the check information message from Host
	for {
		err = client.Call("Host.GetCheckStat", 1, &hostcheckstat)
		if err != nil {
			fmt.Println("call remote service failed", err)
		}
		fmt.Println("remote return success", hostcheckstat)
		if hostmntstat != 0 {
			fmt.Println("get check stat")
			break
		}
		time.Sleep(3 * time.Second)
	}
	//TODO: compare check stat and return test result
	container.IsResultGet = 1
	fmt.Println("get result")
	// CLOSE server and exit
	l.Close()

}
