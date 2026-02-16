package slice_eg

import (
	"fmt"
	"slices"
)

func Slices() {
	
	var s []string 
	fmt.Println("s:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("s:", s, len(s), cap(s))

	
	var s2 []string 
	s2 = make([]string, 0, 3)
	fmt.Println("s:", s2, len(s2), cap(s2))

	// can directly assign because of prefilled zero value
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("s:", s, len(s), cap(s))

	// need to append because of non prefilled zero value
	s2 = append(s2, "a")
	s2 = append(s2, "b")
	s2 = append(s2, "c")
	fmt.Println("s:", s2, len(s2), cap(s2))


	str1 := make([]string, len(s2))
	str2 := []string{"a2","b2","c2"}
	copy(str1, str2)
	str1 = append(str1, "d2")
	str1 = append(str1, "e2")
	str1 = append(str1, "f2")
	str1 = append(str1, "g2")
	fmt.Println("str1:", str1)

	fmt.Println("low:high", str1[2:5])
	fmt.Println("high:", str1[:5])
	fmt.Println("low:", str1[2:])
	

	if slices.Equal(str1, str2) { 
		fmt.Println("same")
	}

	
	return
}
