package main

import "fmt"

func main() {
	//var a, b int = 10, 30
	//fmt.Println(a + b)
	//c,d,e  := numbers()
	//fmt.Println(c,d,e)
	_,result := funcAdd(1,3)
	fmt.Println(result)
}

func numbers() (int,int,string) {
	a,b,c  := 1,3,"hello gogo"
	return a,b,c
}

func funcAdd(a int,b int) (int,int) {
	return  a + b,2*(a + b)
}