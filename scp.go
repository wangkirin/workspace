package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/ssh"
)

const (
	user     = "root"
	ip_port  = "10.229.40.140:22"
	password = "wang"
)

func main() {
	PassWd := []ssh.AuthMethod{ssh.Password(password)}
	Conf := ssh.ClientConfig{User: user, Auth: PassWd}
	Client, err := ssh.Dial("tcp", ip_port, &Conf)
	if err != nil {
		fmt.Println(nil)
	}
	defer Client.Close()
	if session, err := Client.NewSession(); err == nil {
		defer session.Close()
		go func() {
			Buf := make([]byte, 1024)
			w, _ := session.StdinPipe()
			defer w.Close()
			File, _ := os.Open("/root/scptest/test.txt")
			info, _ := File.Stat()
			fmt.Fprintln(w, "C0644", info.Size(), "/root/test.txt")
			for {
				n, err := File.Read(Buf)
				fmt.Fprint(w, string(Buf[:n]))
				if err != nil {
					if err == io.EOF {
						return
					} else {
						panic(err)
					}
				}
			}
		}()
		if err := session.Run("/usr/bin/scp -qrt /mnt"); err != nil {
			if err != nil {
				if err.Error() != "Process exited with: 1. Reason was:  ()" {
					fmt.Println(err.Error())
				}
			}
		}
	}
}
