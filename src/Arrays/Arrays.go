package main
import (
         "fmt"
         "time"
       )
func main () {
	var x [5] int
	x[4] = 100
	fmt.Println(x)
        var notes [7]string
        notes[0] = "do"
        notes[1] = "re"
        notes[2] = "mi"
        notes[3] = "fa"
        notes[4] = "sol"
        notes[5] = "la"
        notes[6] = "ti"
        fmt.Println(notes[0])
        fmt.Println(notes[1])
        fmt.Println(notes[2])
        fmt.Println(notes[3])
        fmt.Println(notes[4])
        fmt.Println(notes[5])
        fmt.Println(notes[6])
        var primes [5]int
        primes[0] = 2
        primes[1] = 3
        fmt.Println(primes[0])
        var dates [3]time.Time
        dates[0] = time.Unix(1257894000, 0)
        dates[1] = time.Unix(1447920000, 0)
        dates[2] = time.Unix(1508632200, 0)
        fmt.Println(dates[1])
}
