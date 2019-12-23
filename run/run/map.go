package main

import "fmt"
import "github.com/skip2/go-qrcode"

func makemap()  {
	maps := make(map[string]int)
	maps = map[string]int{"age":33}
	value,isok := maps["age"]
	delete(maps,"age")
	fmt.Println(value,isok)
	fmt.Println(maps)
}
func main() {
	qrcode.WriteFile("https://www.jianshu.com/p/cc1ffa5a3f4d", qrcode.Medium, 256, "./golang_qrcode.png")

	makemap()
	//var firstMap map[string]int
	////firstMap["key1"]  = 1
	//firstMap = map[string]int{"name": 18, "message": 99}
	//fmt.Println(firstMap)
	//for item, cvalue := range firstMap {
	//	fmt.Println(item,cvalue,"\n")
	//}
	//mapArr()
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