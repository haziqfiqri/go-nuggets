package pointers

import "fmt"

func ZeroValue(n int) {
	n = 0
}

func ZeroPointer(n *int) {
	*n = 0
}

func Pointers() {
	n := 1
	ZeroValue(n)
	fmt.Println(n)
	ZeroPointer(&n)
	fmt.Println(n)
	
	age := 25
	val := &age
	fmt.Println("Address:", age)
	fmt.Println("Address:", val)
	fmt.Println("Value:", *val)

	return	
}
