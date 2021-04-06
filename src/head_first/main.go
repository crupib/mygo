package main
import "fmt"
func main() { 
		var counters [3] int
		var initcount = [3] int {9,18,27}
		var notes[7]string = [7]string{"do", "re", "mi","fa", "so","la","ti"}
		primes := [5] int {2,3,5,7,11}
		counters[0]++
		counters[0]++
		counters[2]++
		index := 3
		fmt.Println(index, notes[index])
		index = 1
		fmt.Println(index, notes[index])
		fmt.Println(counters [0], counters[1], counters[2])
		fmt.Println(initcount [0], initcount[1], initcount[2])
		fmt.Println(notes)
		fmt.Println(primes)
		for i := 0; i <= 6; i++ {
			fmt.Println(i, notes[i])
		}
		
}
