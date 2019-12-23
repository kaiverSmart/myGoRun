package main

import (
	"fmt"
)

//数据从一个线程传到你一个线程
func main() {
	//var ch chan string
	//ch := make(chan string)
	//go sendData(ch)
 	//go getData(ch)
 	//suck(pump())
 	//ch0 := make(chan int)
	//ch1 := make(chan int)
 	//go inputfirst(ch0)
	//go inputsecend(ch1)
	//go out(ch0,ch1)

	//time.Sleep(1e9)
	resume = integers()
	fmt.Println(generateInteger())  //=> 0
	fmt.Println(generateInteger())  //=> 1

}

var resume chan int

func integers() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			yield <- count
			count++
		}
	}()
	return yield
}

func generateInteger() int {
	return <-resume
}



//准备好了吗 可以执行
func inputfirst (ch chan int) {
	for i := 0;;i++ {
		ch <- i * 2
	}
}

func inputsecend(ch chan int) {
	for i := 2; ;i++  {
		ch <- i + 100
	}
}

func out(ch0,ch1 chan int) {
	for {
		select {
		case  v0 := <- ch0 :
		fmt.Println("chan 0",v0)
		case  v1 := <- ch1 :
			fmt.Println("chan 1",v1)
		default:
			fmt.Println(".......")
		}
	}
}
func sendData(ch chan string) {
	ch <- "gz"
	ch <- "sz"
}

func getData (ch chan string) {
	for {
		inputStr := <-ch
		fmt.Println(inputStr)
	}
}



func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}