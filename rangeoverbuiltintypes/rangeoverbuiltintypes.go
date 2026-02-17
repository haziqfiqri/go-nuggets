package rangeoverbuiltintypes

import "fmt"

func RangeOverBuiltInTypes() {

	employeesSalary := make(map[string]int)

	employeesSalary["Ali"] = 1000
	employeesSalary["Budi"] = 2000
	employeesSalary["Cindy"] = 3000

	totalEmployeesSalary := 0
	employeeT20 := make(map[string]int)

	for name, salary := range employeesSalary {
		totalEmployeesSalary += salary

		if salary >= 2000 {
			employeeT20[name] = salary
		}
	}

	tempKeys := employeesSalary

	for key, value := range tempKeys {
		fmt.Println("key:", key)
		fmt.Println("value:", value)
	}

	fmt.Println("totalEmployeesSalary:", totalEmployeesSalary)
	fmt.Println("employeeT20:", employeeT20)
}