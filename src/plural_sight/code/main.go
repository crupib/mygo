package main

import (
	"net/http"
	"plural_sight/code/controllers"
)
func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}