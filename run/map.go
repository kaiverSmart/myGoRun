package main

import "fmt"

func main() {
	var firstMap map[string]int
	//firstMap["key1"]  = 1
	firstMap = map[string]int{"name": 18, "message": 99}
	fmt.Println(firstMap)
	for item, cvalue := range firstMap {
		fmt.Println(item,cvalue,"\n")
	}
	mapArr()
}

func mapArr() {

	var items0 = make([]map[string]int, 5)
	for i := range items0 {
		items0[i] = make(map[string]int)
		items0[i]["title"] = 2
	}

	fmt.Println("item0",items0)

   //items1 := make([]map[string]int,5)
	//for  _,item := range items1 {
	//	item = make(map[string]int,1)
	//	item["titile"] = 2
	//}
   //
	//fmt.Println("item1",items1)
}