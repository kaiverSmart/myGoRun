package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//var isUrl = strings.HasPrefix("http:papa.com","http")
	//
	//fmt.Printf("isUrl:%d",isUrl)

	var numStr = "qe"

	ruleNum, err := strconv.Atoi(numStr)

	fmt.Println("%d", ruleNum, err)

	var t = time.Now()

	var tStr = t.Format("2003:12:13")
	fmt.Println("tstr:", tStr, "t:\n", t)

	i := 0

HEREE:
	fmt.Println(i)
	i += 1
	if i == 5 {
		return
	}
	goto HEREE
}
