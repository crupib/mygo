package main

import "fmt"

func IncrementPassByPointer(ptr *int) {
   (*ptr)++
}
func main() {
   i:=10
   fmt.Println("i = ::",i)
   IncrementPassByPointer(&i)
   fmt.Println("i = ::",i)
}
