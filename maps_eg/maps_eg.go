package maps_eg

import (
	"fmt"
	"maps"
)

func Maps() {
	
	malaysiaStates := make(map[string]int);
	fmt.Println("malaysiaStates:", malaysiaStates)

	malaysiaStates["selangor"] = 1
	malaysiaStates["wilayah_persekutuan"] = 2
	malaysiaStates["melaka"] = 3

	fmt.Println("malaysiaStates:", malaysiaStates)

	
	value := malaysiaStates["selangor"]

	fmt.Println("value:", value)
	fmt.Println("len:", len(malaysiaStates))

	_, ok := malaysiaStates["selangor"]
	val, ok := malaysiaStates["selangor"]

	fmt.Println("ok:", ok)
	fmt.Println("val:", val)

	delete(malaysiaStates, "melaka")

	fmt.Println("malaysiaStates:", malaysiaStates)

	fmt.Println("clearing...")
	clear(malaysiaStates)

	fmt.Println("malaysiaStates:", malaysiaStates)

	_, ok = malaysiaStates["selangor"]

	fmt.Println("ok:", ok)

	selangorPostcode := map[string]int{"4010": 1, "4015": 2, "4020": 3}

	fmt.Println("selangorPostcode:", selangorPostcode)

	userPostcode := map[string]int{"4010": 1}

	if maps.Equal(userPostcode, selangorPostcode) {
		fmt.Println("same")
	} else {
		fmt.Println("not same")
	}

	return
}
