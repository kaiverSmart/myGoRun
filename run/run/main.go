package main

import (
   "fmt"
   "log"
)

func main() {
   fmt.Println("Hello, World!")
   var str  = "string0"
   fmt.Printf(str)

   fmt.Println("end program")
   //panic("a server error")
//protect(main)

}

func protect(g func()) {
   defer func() {
      log.Println("done")
      // Println executes normally even if there is a panic
      if err := recover(); err != nil {
         log.Printf("run time panic: %v", err)
      }
   }()
   log.Println("start")
   g() //   possible runtime-error
}