package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main()  {
	//inputFile,inputError := os.Open("/Users/g/Desktop/GOGO/myGoRun/run/test.m")
	//if inputError != nil {
	//	fmt.Println(inputError)
	//	return
	//}
	//
	//defer  inputFile.Close()
	//inputReder := bufio.NewReader(inputFile)
	//for  {
	//	inpStr, error := inputReder.ReadString('\n')
	//	fmt.Println(inpStr)
	//	if error != nil {
	//		fmt.Println(error)
	//		return
	//	}
	//}
	fileStr := "/Users/g/Desktop/GOGO/myGoRun/run/test.m"
	copyStr := "/Users/g/Desktop/GOGO/myGoRun/run/copy.m"
	start := time.Now()
	 wirtten,error := copyFile(fileStr,copyStr)

	if error != nil {
		fmt.Println(error)
	}
	end := time.Now()
	detailTime := end.Sub(start)
	fmt.Println(wirtten,detailTime)
}

func copyFile(fileName string,copyFileName  string) (written int64, err error)  {
	currentFile,fileErr := os.Open(fileName)
	if fileErr != nil {
		fmt.Println(fileErr)
	}
	defer currentFile.Close()

	creFile,creErr := os.Create(copyFileName)
	if creErr != nil {
		fmt.Println(creErr)
	}

	return io.Copy(creFile,currentFile)
}