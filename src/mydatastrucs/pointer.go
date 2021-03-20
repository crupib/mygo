package main

import "fmt"

func main() {
   data := 10
   ptr := &data
   fmt.Println("value stored at variable var is",data)
   fmt.Println("value stored at variable var is",*ptr)
   fmt.Println("address stored at variable var is",&data)
   fmt.Println("address stored at variable var is",ptr)
}

