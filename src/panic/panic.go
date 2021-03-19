package main

import (
    "fmt"
)

func main() {
	panic("PANIC")
        defer func() {
		str := recover()
		fmt.Println(str)
	} ()
	panic("PANIC")
}
