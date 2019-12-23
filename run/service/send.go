package main

import (
	"fmt"
	"time"
)

var valuse = []int {1,2,3,4}
func main() {
	for id := range valuse  {
	 func() {
		fmt.Println("idx:",id)
		//fmt.Println("iv:",iv)
	 }()
	}

	for id := range valuse  {
		go func() {
			fmt.Println("idx:",id)
			//fmt.Println("iv:",iv)
		}()
	}

	time.Sleep(1e9)
}