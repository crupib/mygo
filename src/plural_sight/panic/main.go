package main
func main() {
	println("Before panic")
	panic("I am fucked")
	println("after panic will never get here!")
}
