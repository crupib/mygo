package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	fmt.Println("Switch statement")
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1
	fmt.Println("Day",dow)
	var result string
	switch dow {
	case 1:
		result = "It's Sunday!"
		fallthrough
	case 2:
		result = "It's Monday!"
	case 3:
		result = "It's Tuesday!"
	case 4:
		result = "It's Wednsday!"
	case 5:
		result = "It's Thursday!"
	case 6:
		result = "It's Friday!"
	case 7:
		result = "It's Saturday!"
	default:
		result = "Some type of error"
	}
	fmt.Println(result)
}