package main

import "fmt"

func main() {
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	fmt.Println(elements["Li"])
	fmt.Println(elements["Un"])
	name, ok := elements["B"]
	fmt.Println(name, ok)
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}
}
