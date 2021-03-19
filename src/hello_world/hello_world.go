package main
import (
	"fmt"
	"runtime"
)
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
func makeEvenGenerator() func() uint {
	h := uint(0)
	return func() (ret uint) {
		ret = h
		h += 2
		return
	}
}
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
func f() {

	y := 1
	fmt.Println("Hello, World!",len("Hello, World!"))
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
func average(xs []float64) float64 {
	total :=0.0
	for _,v := range xs {
		total += v
	}
	if len(xs) == 0 {
	   panic ("Cannot divede by zero")}
	return total/float64(len(xs))
}
func f2() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["H"] = "Hydrogen"
	fmt.Println(elements["He"])
}
var x string = "Shit for brains"
const z string = "hello world"
func main() {
	var ( a = 5
	      b = 10
	      c=15
	      )
	fmt.Println(a, b, c)
	f()
	fmt.Print("Enter a number: ")
//	var input float64
//	fmt.Scanf("%f",&input)
//	output := input * 2
//	fmt.Println(output)
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	var w [5]int
	w[4] = 1000
	fmt.Println(w)

	t := [5]float64{1,2,3,4,5}
	fmt.Println(t, len(t))
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4,5)
	fmt.Println(slice1, slice2)
	slice3 := make([]int,5)
	copy(slice3,slice2)
	fmt.Println(slice3)
	f2()
	xs := []float64{99,98,93,22,93}
//	var xt []float64
	fmt.Println(average(xs))
//	fmt.Println(average(xt))
	fmt.Println(add(1,2,3))
	fmt.Println(add(1,2))
	fmt.Println(add(1,2,3,4,5))
	add_inline := func(x,y int) int {
		return x + y
	}
	fmt.Println(add_inline(95,1))
	j := 0
	increment := func() int {
		j++
		return j
	}
	fmt.Println(increment())
	fmt.Println(increment())
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(factorial(5))
}
