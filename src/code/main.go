package main

import (
	"code/models"
	"fmt"
)

func main() {

     u := models.User {
          ID: 2,
          FirstName: "Joe",
          LastName: "Blow",
     }
     fmt.Println(u)

}
