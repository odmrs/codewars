package main

func main() {
	Solution("Marcos")
}

func Solution(word string) string {
	var reversed string
	for i := 1; i < len(word)+1; i++ {
		reversed += string(word[len(word)-i])
	}
	return reversed
}
