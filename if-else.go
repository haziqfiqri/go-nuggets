package main

import "fmt"

func main() {

	const result = 1 + 1

	if result % 2 == 0 {
		fmt.Println("result is even")
	} else {
		fmt.Println("result is odd")
	}

	const user_conn = true

	if user_conn {
		fmt.Println("user is connected")
	}

	if user_conn := false; user_conn {
		fmt.Println("user is connected")
	} else {
		fmt.Println("user is not connected")
	}

	if userAge := 11; userAge >= 18 {
		fmt.Println("adult")
	} else if userAge >= 13 {
		fmt.Println("teenager")
	} else if userAge >= 1 {
		fmt.Println("child")
	} else {
		fmt.Println("unborn")
	}
}
