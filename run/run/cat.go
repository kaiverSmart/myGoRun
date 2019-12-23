package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
	return
}
func cat2(f *os.File) {
	const NBUF = 512
	var buf [NBUF]byte
	for {
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0: // EOF
			return
		case nr > 0:
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew.Error())
			}
		}
	}
}

func readData() string{
	f,err := os.OpenFile("",os.O_RDWR,0)
	if err != nil{
		fmt.Println(err)
		return err.Error()
	}
	ioutil.ReadAll(f)
}
func main () {
	fileStr := "/Users/g/Desktop/GOGO/myGoRun/run/test.m"
	f,err := os.Open(fileStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	cat2(f)
}
//func main() {
//	flag.Parse()
//	if flag.NArg() == 0 {
//		cat(bufio.NewReader(os.Stdin))
//	}
//	for i := 0; i < flag.NArg(); i++ {
//		f, err := os.Open(flag.Arg(i))
//		if err != nil {
//			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
//			continue
//		}
//		cat(bufio.NewReader(f))
//	}
//}
