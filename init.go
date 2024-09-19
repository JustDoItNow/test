package main

import "fmt"

var count int

func init() {
	count = -100
	fmt.Println("init() function is called: count is initilized to ", count)
}
