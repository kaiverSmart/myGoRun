package main

import "fmt"

type person struct {
	name string
	age int
}
func main ()  {
	var p1 = person{"rose",19}
	fmt.Println(p1)

 	fmt.Printf("name:%s age:%d \n",p1.name,p1.age)

 	var p2 = person{age:10}

 	fmt.Printf("name:%s age:%d",p2.name,p2.age)
}