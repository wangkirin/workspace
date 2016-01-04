package main

import (
	"fmt"
	"os"
)

func main() {
	Startservers()
}

func Startservers() error {
	fmt.Println("===============enter start servers")
	os.Getwd()
	binpath := "./oss/chunkserver/spy_server"
	errlogpath := "/root/gopath/chunkserver/errlog/err.log"
	// check if chunkserver binary exsist
	_, err := os.Stat(binpath)
	if err != nil && os.IsNotExist(err) {
		return fmt.Errorf("Cannot find chunkserver excution file")
	}
	fmt.Println("===============check binary done")
	// check if errlog folder exsist , if not ,create it
	_, err = os.Stat(errlogpath)
	if err != nil || os.IsNotExist(err) {
		_, err := os.OpenFile(errlogpath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("===============check errloir done")
	fmt.Printf("===============len cs=%v \n", len(this.cs))
	for i := 0; i < len(this.cs); i++ {
		go func() {
			fmt.Println("===============enter goroutine")
			fmt.Printf("===============i=%v \n", i)
			curcs := this.cs[i]
			_, err := os.Stat(curcs.DataDir)
			if err != nil || os.IsNotExist(err) {
				os.MkdirAll(curcs.DataDir, 0777)
			}
			Port := fmt.Sprintf("%v", curcs.Port)
			fmt.Println(errlogpath)
			out, err := exec.Command("./oss/chunkserver/spy_server", "--ip", curcs.Ip, "--port", Port, "--master_ip", string(this.cm.serverHost), "--master_port", "8099", "--group_id", "1", "--chunks", "2", "--data_dir", curcs.DataDir, "--error_log", errlogpath).CombinedOutput()
			fmt.Println("out==" + string(out))
			if err != nil {
				fmt.Println(err.Error())
			}
		}()
		runtime.Gosched()
	}
	return nil
}
