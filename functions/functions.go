package functions

import "fmt"

func Addition(a int, b int) int {
	return a + b
}

func GetUser(name string, age, phone int) (string, int, int) {
	return name, age, phone
}

func DbAdapter(name, dialect string, port int) (string, string, int) {
	return name, dialect, port
}

func VariadicFunc(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func VariadicFuncWithOtherParams(a int, b string, nums ...int) (int, string) {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total, b
}

func ClosureFunc() func() int { 
	return func () int { 
		return 1
	}
}

func ClosureFunc2() func() int {
	i := 0

	return func() int { 
		i++
		return i
	}
}

func RecursionFunc(n int) { 
	if n > 0 { 
		fmt.Println(n)
		RecursionFunc(n - 1)
	} else {
		fmt.Println("countdown stop")
	}
}