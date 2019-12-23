package main

import "fmt"

func main ()  {
	var num = 20
	var ip  = &num


	fmt.Println("变量值",*ip)

	*ip =  30

	fmt.Println("num",num)


	fmt.Println("指针变量%x",&ip)

	fmt.Println("指针变量%x",&num)
}