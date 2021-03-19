package main

import "fmt"

var c, python, java bool
var i, j int = 1, 2
func main() {
	fmt.Println(i, c, python, java)
        k := 3
        var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	fmt.Println(i, j, k, c, python, java)
}

