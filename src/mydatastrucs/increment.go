package main

import "fmt"

func IncrementPassByValue(x int) {
   x++
   fmt.Println("x = ::",x)
}
func main() {
   x:=10
   fmt.Println("x = ::",x)
   IncrementPassByValue(x)
   fmt.Println("x = ::",x)
}
