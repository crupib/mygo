package main

import "fmt"
import "gopkg.in/yaml.v2"

type Person struct {
    Name string `json:"Name" yaml:"Name"` // Supporting both JSON and YAML.
    Age int `json:"age" yaml:"age"`
}
func main() {
    p := Person{"John", 30}
// Convert the Person struct to YAML.
    y, err := yaml.Marshal(p)
    if err != nil {
       fmt.Println(err)
       return 
    } 
    fmt.Println(string(y))
}
