package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	var pwd string
	var host string
	var port int

	flag.StringVar(&name, "n", "admin", "输入用户名")
	flag.StringVar(&pwd, "pwd", "", "输入密码")
	flag.StringVar(&host, "h", "localhost", "输入主机名")
	flag.IntVar(&port, "port", 8888, "输入端口")

	flag.Parse()

	fmt.Printf("name=%v,pwd=%v,host=%v,port=%v", name, pwd, host, port)

}
