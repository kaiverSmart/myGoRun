package main

import (
	"../pack1"
	"fmt"
	"net"
)

//终端输入发起网络请求

func main() {
	pack1.Pack1()
 	con,err :=	net.Dial("tcp","localhost:50000")
	if err != nil {
		fmt.Println("con",err)
	}
	buffStr := []byte("send mes")
	con.Write(buffStr)
}