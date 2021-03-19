package main
import "fmt"
import "os"
func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}
func main() {
     defer second()
     first()
     f, _ := os.Open("/etc/hosts")
     defer f.Close()
}
