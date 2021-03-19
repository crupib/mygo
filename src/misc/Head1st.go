package main

import (
	"math"
	"strings"
	//     "net/http"
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"
)

func main() {
	quantity := 4
	length, width := 1.2, 2.4
	customerName := "Damon Cole"
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(year)
	fmt.Println("ㅇ솓 아")
	fmt.Println(now)
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("head \tfirst\t go"))
	fmt.Println("\n\nCannonball!!!!\n")
	fmt.Println('A')
	fmt.Println('Á', 1, 1.234)
	fmt.Println(reflect.TypeOf(3.1415))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello, Go!"))
	fmt.Println(reflect.TypeOf(42424242))
	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")
	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
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
	//fmt.Println(input)
}
