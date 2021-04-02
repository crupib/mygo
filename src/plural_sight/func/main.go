package main

import (
	"errors"
	"fmt"
)
func main() {
	fmt.Println("Hello, playground")
	port := 3000
	p, err := StartWebServer(port,2)
	fmt.Println(p, err)
}
func StartWebServer(port int, numberOfRetries int) (int,error) {
	fmt.Println("Starting server...")
	fmt.Println("Server started on port",port)
	fmt.Println("Number of retries", numberOfRetries)
	return port, errors.New("Something fucked up")
}