package main

import (
       "log"
       "util"
       "fmt"
)
func main() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := util.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
        celsius := (fahrenheit - 32) * 5/9
	fmt.Println("%0.2f degrees Celsius\n", celsius)
}
