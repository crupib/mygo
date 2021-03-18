package main
import (
          "util"
          "fmt"
          "log"
       )
func main() {
    fmt.Print("Enter a float: ")
    myfloat, err  := util.GetFloat()
    if err != nil {
       log.Fatal(err)
    }
    fmt.Println(myfloat)
    fmt.Println(myfloat)
}
