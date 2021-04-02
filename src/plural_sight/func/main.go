package main

import (
	"errors"
	"fmt"
)
func main() {
	fmt.Println("Hello, playground")
	port := 3000
	err := StartWebServer(port,2)
	fmt.Println(err)
}
func StartWebServer(port int, numberOfRetries int) error {
	fmt.Println("Starting server...")
	fmt.Println("Server started on port",port)
	fmt.Println("Number of retries", numberOfRetries)
	return errors.New("Something fucked up")
}