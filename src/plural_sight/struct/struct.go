package main

import (
	"fmt"
)

func main() {
	type user struct {
		ID int
		FirstName string
		LastName string
	}
	var u user
    u.ID = 1
	u.FirstName = "Bill"
	u.LastName = "Crupi"
    fmt.Println(u)

}
