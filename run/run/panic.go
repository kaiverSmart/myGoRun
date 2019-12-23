package main

import (
	"flag"
	"fmt"
	"runtime"
)

func main() {
	//f()
	var numCore = flag.Int("namef",4,"usages")
	runtime.GOMAXPROCS(*numCore)
}

func f()  {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover",r)
		}
	}()
	fmt.Println("run f")
	g(0)
	fmt.Println("run f over")
}
func g(i int)  {
	fmt.Println("run g",i)
	if i > 3 {
		fmt.Println("current index",i)
		panic(fmt.Sprint("run g panic ",i))
	}
	g(i + 1)
}