package main

import "fmt"
import "math"
const PI = 3.14
func max(x,y int) int {
   var max int
   if x > y {
     max = x
   } else {
     max = y
   }
   return max
}
func maxAreaCheck(length, width, limit int) bool {
   if area := length*width; area < limit {
        return true
   } else {
        return false
   }
}
func main() {
   var v1 int
   v2 := 200
   v1 = 100
   fmt.Println("Value stored in variable v1:: ",v1)
   fmt.Println("Value stored in variable v2:: ",v2)
   fmt.Println("Pi = :: ",PI)
   fmt.Println("Hello, World!")
   fmt.Println("Range of Int8::",math.MinInt8,"to",math.MaxInt8)
   s := "hello, World!"
   r := []rune(s)
   r [0] = 'F'
   r [1] = 'U'
   r [2] = 'C'
   r [3] = 'K'
   s2 := string(r)
   fmt.Println(s2)
   fmt.Println(max(30,21))
   fmt.Println("area check is", maxAreaCheck(20,10,300))
   i:= 2
   switch i {
   case 1:
      fmt.Println("One")
   case 2:
      fmt.Println("Two")
   case 3:
      fmt.Println("Three")
   default:
      fmt.Println("Sonething else")
   }
}
