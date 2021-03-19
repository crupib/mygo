package main

import "fmt"
func max(x, y int) int {
   if x > y {
      return x
   }
   return y
}
func main() {
   fmt.Println("max of 2 numbers :: ",max(40,30))
}

