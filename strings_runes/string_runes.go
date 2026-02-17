package strings_runes

import (
	"fmt"
	"unicode/utf8"
)

func StringRunes() {
	var word string = "電車"

	fmt.Println(len(word))
	fmt.Println(utf8.RuneCountInString(word))

	for i := 0; i < len(word); i++ { // it treat like a slice of byte
		fmt.Println("for", word[i])
	}

	for idx, runeValue := range word { // it transform byte into runes
		fmt.Println("range", idx, runeValue)
	}

	// to achieve same concept like range which is using decoderuneinstring under the hood
	for i, rs := 0, 0; i < len(word); i += rs {
		runeValue, runeSize := utf8.DecodeRuneInString(word[i:])

		fmt.Println(runeValue, runeSize)

		rs = runeSize

		ExamineRune(runeValue)
	}

	return
}

func ExamineRune(runeValue rune) {
	switch runeValue {
		case 't':
			fmt.Println("found t")
		case '電':
			fmt.Println("found 電")
		default:
			fmt.Println("not found")
	}
}