package main
import "fmt"
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total/float64(len(xs))
}
func f() (int, int) {
	return 5,6
}
func main() {
        add1 := func(x, y int) int{ return x+y
		}
        fmt.Println(add1(1,1))	
	xs := []float64{98,93,77,82,83}
	fmt.Println(average(xs))
        x, y := f()
        fmt.Println(x, y)
        fmt.Println(add(1,2,3,4,5,6,7,8,9,10))
        xs_1 := []int{1,2,3}
	fmt.Println(add(xs_1...))
        x = 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
