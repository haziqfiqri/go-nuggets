package arrays

import "fmt"

func Arrays() { 
	var arr [5]string 
	fmt.Println("emp:", arr)

	fmt.Println("len:", len(arr))

	arr[4] = "hassan"
	fmt.Println("set:", arr)
	fmt.Println("get:", arr[4])

	fmt.Println("len:", len(arr))

	newarr := [5]int{1,2,3,4,5}
	fmt.Println("newarr:", newarr)

	// This works because the highest index reached is 4 (Total size: 5)
	// Index 0: 100
	// Index 1: 0 (auto-filled)
	// Index 2: 0 (auto-filled)
	// Index 3: 400
	// Index 4: 500
	newarr = [...]int{100, 3: 400, 500}
	fmt.Println("working sparse example:", newarr)

	// This failed because Index 5 was reached (Total size: 6)
	// Index 0: 1
	// Index 1: 0 (auto-filled)
	// Index 2: 200
	// Index 3: 3
	// Index 4: 4
	// Index 5: 5	
	sparseArr := [...]int{1, 2: 200, 3, 4, 5}
	fmt.Println("sparseArr (size 6):", sparseArr)

	// Array default are 1D
	one_d_arr := [5]int{1,2,3,4,5}
	fmt.Println("one_d_arr:", one_d_arr)

	// Multi-dimensional array
	two_d_arr := [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println("two_d_arr:", two_d_arr)

	var two_d_arr_2[2][3]int

	for i := range 2 { 
		for j := range 3 { 
			two_d_arr_2[i][j] = i + j
		}
	}
	fmt.Println("two_d_arr_2:", two_d_arr_2)

	return
}