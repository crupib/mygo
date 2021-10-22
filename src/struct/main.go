package main

import (
	"fmt"
)
func main() {
//	fmt.Println("Structures")
	poodle := Dog{"Poodle", 10, "Woof!"}
//	fmt.Println(poodle)
//	fmt.Printf("%+v\n",poodle)
//	fmt.Printf("Breed: %v\nWeight: %v\nSound: %v\n", poodle.Breed, poodle.Weight, poodle.Sound)
	poodle.Weight = 9
	fmt.Printf("Breed: %v\nWeight: %v\nSound: %v\n", poodle.Breed, poodle.Weight, poodle.Sound)
	poodle.Speak()
	poodle.Sound = "Arf!"
	poodle.SpeakThreeTimes()
}
// Dog is a struct
type Dog struct {
	Breed string
	Weight int
	Sound string
}
// Speak is how the dog speaks
func (d Dog) Speak(){
	fmt.Println(d.Sound)
}
// SpeakThreeTimes Speak three times
func (d Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}