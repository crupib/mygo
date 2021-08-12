package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)

	fmt.Println("Maps")
	fmt.Println(states)
	states["WA"] = "Washington"
	states["FL"] = "Florida"
	states["NJ"] = "New Jersey"
	states["NY"] = "New York"
	states["CA"] = "California"
	fmt.Println(states)
    california := states["CA"]
	fmt.Println(california)
	delete(states,"WA")
	states["MI"] = "Missouri"
	fmt.Println(states)
	for k, v := range states {
		fmt.Printf("%v : %v \n",k,v)
	}
	keys := make([]string, len(states))
	i:=0
	for k := range states{
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)
	for i:= range keys{
		fmt.Println(states[keys[i]])
	}
}