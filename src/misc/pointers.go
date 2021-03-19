package main

import (
    "fmt"
    "reflect"
)
func createPointer() *float64 {
   var myFloat = 98.5
   return &myFloat
}
func printPointer(myBoolPointer *bool) {
   fmt.Println(*myBoolPointer)
}

func main() {
   var myInt int
   var myIntPointer *int
   fmt.Println(&myInt)
   fmt.Println(reflect.TypeOf(&myInt))
   var myFloat float64
   fmt.Println(&myFloat)
   fmt.Println(reflect.TypeOf(&myFloat))
   var myBool bool
   fmt.Println(&myBool)
   fmt.Println(reflect.TypeOf(&myBool))
   amount := 6
   fmt.Println(amount, &amount)
   myIntPointer = &myInt
   fmt.Println(myIntPointer)
   var myFloatPointer *float64 = createPointer()
   fmt.Println(*myFloatPointer)
   myBool = true
   printPointer(&myBool)
}
   
