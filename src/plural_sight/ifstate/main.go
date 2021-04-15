package main


type User struct {
	ID 	int
	FirstName string
	LastName string
}
func main() {
	u1 := User{
		ID: 1,
		FirstName: "Arther",
		LastName: "Dent",
	}
	u2 := User {
		ID: 2,
		FirstName: "Arther",
		LastName: "Prefect",
	}
	if u1.ID == u2.ID {
		println("Same user!")
	} else  if u1.FirstName == u2.FirstName {
		println("Similar users!")
	} else {
		println("Different user")
	}
}