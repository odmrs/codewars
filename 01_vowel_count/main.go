package main

import "fmt"

func main() {
	fmt.Println(GetCount("aeiou"))
}

var vowels []string = []string{"a", "e", "i", "o", "u"}

func GetCount(str string) (count int) {
	// for _, letter := range str {
	// 	for i := 0; i < len(vowels); i++ {
	// 		if string(letter) == vowels[i] {
	// 			count += 1
	// 		}
	// 	}
	// }

	// Or more clean

	for _, letter := range str {
		switch letter {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}

	return
}
