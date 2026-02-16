package switches

import (
	"fmt"
	"time"
)

func Switches() { 

	i := 2
	fmt.Println("Write", i, "as")

	switch i { 
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		default:
			fmt.Println("other")
	}

	fmt.Println("What time is today")

	var now = time.Now()

	switch now.Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("weekend")
		default:
			fmt.Println("weekday")
	}

	fmt.Println("Role base access control")

	var role = "admin"

	switch role {
		case "admin":
			fmt.Println("full access")
		case "user":
			fmt.Println("limited access")
		case "guest":
			fmt.Println("partial access")
		default:
			fmt.Println("unauthorized access")
	}

	fmt.Println("No expression is treated like if/else")

	var t = time.Now()

	switch {
		case t.Hour() < 12:
			fmt.Println("Before afternoon")
		default:
			fmt.Println("Its afternoon")
	}

	fmt.Println("A type switch compares type instead of values")

	whatami := func(i interface{}) { //or func(i any) 
		switch t := i.(type) { 
			case bool:
				fmt.Println("I'm a bool")
			case int:
				fmt.Println("I'm an int")
			default:
				fmt.Println("I'm something else \n", t)
		}
	}

	whatami(true)
	whatami(1)
	whatami("hello")
}