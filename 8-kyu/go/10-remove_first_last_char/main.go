package main

import "fmt"

func main() {
	var splitedString = RemoveChar("Marcos")
	fmt.Println(splitedString)
}

func RemoveChar(word string) string {
	// return word[1 : len(word)-1]
	// I think the better way is this, because get all utf-8
	var rune_of_word = []rune(word)
	return string(rune_of_word[1 : len(rune_of_word)-1])
}
