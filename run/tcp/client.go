package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
fmt.Println("client start")
conServer()
}
func conServer()  {
	con,err := net.Dial("tcp","localhost:12345")
	defer con.Close()
	if err != nil {
		fmt.Println("con server err",err)
	}
	//监听终端输入
	newRead := bufio.NewReader(os.Stdin)
	for {
		input,err := newRead.ReadString('\n')
		if err != nil {
			fmt.Println("err ",err)
		}
		trimStr := strings.Trim(input,"\r\n")
		if trimStr == "Q" {
			fmt.Println("input end")
			return
		}
		_,writErr := con.Write([]byte(trimStr))
		if writErr != nil {
			fmt.Println("writErr", writErr)
		}
	}
}