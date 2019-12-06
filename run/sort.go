package main

import (
	"fmt"
	"sort"
)

var (
	 currentMap = map[string]int {"name":32,"job":43,"kiss":99,"roes":22}
	)

type Day struct {
	num int
	name string
}

type DayArr struct {
	dataArr  []*Day
}
func (p *DayArr) Len() int           { return len(p.dataArr) }
func (p *DayArr) Less(i, j int) bool { return p.dataArr[i].num < p.dataArr[j].num }
func (p *DayArr) Swap(i, j int)      { p.dataArr[i], p.dataArr[j] = p.dataArr[j], p.dataArr[i] }

func main() {

	d0 :=  Day{0,"sunDay"}
	d1 := Day{1,"monday"}
	d2 := Day{2,"tuesday"}
	//d3 := Day{3,"wednesday"}
	//d4 := Day{4,"thursday"}
	//d5 := Day{5,"friday"}
	//d5 := Day{6,"saturday"}

	days := []*Day {&d1,&d0,&d2}
	a := DayArr{days}
    sort.Sort(&a)

	for  _,d := range days {
		fmt.Println(d)
	}
	//var keys = make([]string,len(currentMap))
	//var values = make([]int,len(currentMap))
	//
	//i := 0
	//for k,v := range currentMap {
	//	fmt.Printf("key:%s value:%d \n",k,v)
	//	keys[i] = k
	//	values[i]  = v
	//
	//	i++
	//}
	//sort.Strings(keys)
	//sort.Ints(values)
	//for _,k := range keys  {
	//	fmt.Printf("key:%s value:%d \n",k,currentMap[k])
	//}
	//
	//for _ , val := range values {
	//	fmt.Printf("value:%d \n",val)
	//}
}