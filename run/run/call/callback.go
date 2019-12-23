package callback

type callS struct {
	name string
}

//func main() {
//	//addCallback(1,addNum)
//	//fmt.Println("interaffece")
//
//	start := time.Now()
//	adderFunc := adder(2)
//	lastNum := adderFunc(3)
//	fmt.Println(lastNum)
//	end := time.Now()
//	deltaTime := end.Sub(start)
//	fmt.Println("time:", deltaTime)
//}
//
//func addNum(a int, b int) {
//	fmt.Println("a+b:", a+b)
//}
//
//func addCallback(y int, f func(int, int)) {
//	f(y, 2)
//}
//
///*函数作为参数传递*/
//func adder(a int) func(b int) int {
//	return func(b int) int {
//		return a + b
//	}
//}
