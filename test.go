package main

// When ... is Required:
// You need ... when you have a slice and want to pass all its elements as individual arguments
func testVariadicFunction() {
	args1 := []interface{}{"example", 42}
	fmt.Printf("This is a message: %s, %d\n", args1)  // This will cause the error (the code can still be built without compile error)
	fmt.Printf("This is a message: %s, %d\n", args1...)
}

// Ouput:
// This is a message: [example %!s(int=42)], %!d(MISSING)
// This is a message: example, 42
