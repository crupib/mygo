package main
import "fmt"

const aConst float32 = 64.0
func main() {
	var aString string = "This is Go!"
	fmt.Println(aString)
	fmt.Printf("The variable's type is %T\n",aString)
	var anInteger int = 42
	fmt.Println(anInteger)
	fmt.Printf("The variable's type is %T\n",anInteger)
	var defaultInt int
	fmt.Println("Default ",defaultInt)
	var anotherString = "This is another string"
	fmt.Println(anotherString)
	fmt.Printf("The variable's type is %T\n",anotherString)
	var anotherInt int = 53
	fmt.Println(anotherInt)
	fmt.Printf("The variable's type is %T\n",anotherInt)
	myString := "This is also a string"
	fmt.Println(myString)
	fmt.Printf("The variable's type is %T\n",myString)
	fmt.Println(aConst)
	fmt.Printf("The variable's of const type is %T\n",aConst)
}