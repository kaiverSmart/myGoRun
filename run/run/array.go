package main

import "fmt"

func main () {
	//var firstArr [3] int

	var firstArr = [3]int{1,4,3}
	var i int
	for i = 0;i < 3 ;i++  {
		fmt.Println(firstArr[i])
	}
}