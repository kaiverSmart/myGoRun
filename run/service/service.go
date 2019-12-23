package main

import (
	"fmt"
	"net"
)

func main(){
	fmt.Println("start service")
	listener,err := net.Listen("tcp","localhost:50000")
	if err != nil {
		fmt.Println("listen err",err)
	}
	// 监听并接受来自客户端的连接
	for {
		con,err_c := listener.Accept()
		if err_c != nil {
			fmt.Println("con err",err_c)
			return
		}
		go precess(con)
	}
}

//网络相应buff
func precess (con net.Conn) {
	var redBuff = make([]byte,512)
	len,err := con.Read(redBuff)
	if err != nil {
		fmt.Println("pres err", err)
	}
	fmt.Println("len ",len)
	str := string(redBuff[:len])
	fmt.Println("receviced \n",str)
}