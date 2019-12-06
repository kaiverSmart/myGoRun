package main

import (
	"fmt"
	"runtime"
)

type person struct {
	name string
	age int
}

type Car struct {
	name string
	logo string
}

type trian struct {
	speed float64
	Car
}
//工厂创建方法
func NewCar (name string,logo string) *Car {
	return &Car{name:name,logo:logo}
}

func main ()  {
	var p1 = person{"rose",19}
	fmt.Println(p1)

 	fmt.Printf("name:%s age:%d \n",p1.name,p1.age)

 	var p2 = person{age:10}

 	fmt.Printf("name:%s age:%d",p2.name,p2.age)


 	c1 := NewCar("toyouta","fengtian")

 	fmt.Println("car:",c1)
	t1 := trian{ 12.2,Car{name: "hexie", logo: "gs"}}
	t1.Car = Car{name: "hexie", logo: "gs"}
	fmt.Println("\n",t1)

	var mem  = runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	fmt.Printf("\nmem:%dkb",mem.Alloc/1024)
}