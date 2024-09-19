package main

import (
  "fmt"
)

// define count in init.go
// var count int

// init() function is a special function that is automatically called by the Go runtime before the main() function. 
// It is used to perform initialization tasks such as setting up global variables, initializing data structures, 
// or any other setup work your application needs before it starts executing.
func init() {
  count = 100
  fmt.Println("init() func is called automatically, count is initialized to ", count)
}

func main() {
  fmt.Println("count value: ", count)
}

// output:
// $ go run .
// init() function is called: count is initilized to  100
// init() function is called: count is initilized to  1
// count value:  1
