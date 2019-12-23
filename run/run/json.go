package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type address struct {
	City string
	cunutry string
}

type vCard struct {
	Name string `json:"name"`
	Adds []address
}

func main() {
	add0 := address{City:"gz",cunutry:"th"}
	addArr := []address{add0}
	vC0 := vCard{Name:"openlee"}
	vC0.Adds = addArr
	fmt.Println(vC0)
   jsonByte,err_json:= json.Marshal(vC0)

	if err_json != nil{
		fmt.Println(err_json)
	}
	jsonStr := string(jsonByte[0:])
	fmt.Println(jsonStr)

	file, _ := os.OpenFile("./vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vC0)
	if err != nil {
		log.Println("Error in encoding json")
	}
}