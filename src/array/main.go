package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors [3]string
	var color_slice = []string{"Red","Green", "Blue"}
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)
	fmt.Println(colors[0])
	var numbers = [5]int{5,3,1,2,4}
	fmt.Println(numbers)
	fmt.Println("Nubers of colors:",len(colors))
	fmt.Println("Nubers of numbers:",len(numbers))
	fmt.Println(color_slice)
    color_slice = append(color_slice,"Purple")
	fmt.Println(color_slice)
	color_slice = append(color_slice[1:len(color_slice)])
	fmt.Println(color_slice)
	color_slice = append(color_slice[:len(color_slice)-2])
	fmt.Println(color_slice)
    numbers2 := make([]int,5,5)
	numbers2[0]= 134
	numbers2[1]= 72
	numbers2[2]= 14
	numbers2[3]= 13
	numbers2[4]= 1
    fmt.Println(numbers2)
	numbers2 = append(numbers2, 235)
	fmt.Println(numbers2)
	sort.Ints(numbers2)
	fmt.Println(numbers2)
}