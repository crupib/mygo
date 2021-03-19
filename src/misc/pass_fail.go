package main

import (
       "log"
       "util"
       "fmt"
)
func main() {
	fmt.Print("Enter a grade: ")
	grade, err := util.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
