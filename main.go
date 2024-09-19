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

// Notes:
// init.go is defined in the same package (it doesn't necessarily to be named 'init.go', just to highligh its functionality)
// main.go is loaded first by runtime, so its init() is called first, and then the init() func in init.go

// output:
// $ go run .
// init() function is called: count is initilized to  100
// init() function is called: count is initilized to  1
// count value:  1
