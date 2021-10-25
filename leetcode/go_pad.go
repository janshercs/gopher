package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("hi")
}

func driver(input string) bool {
	return isPalindrome(input)
}
func isPalindrome(s string) bool {
	r, err := regexp.Compile("[^a-z0-9]")
	if err != nil {
		log.Fatal(err)
	}
	s = strings.ToLower(s)
	processedString := r.ReplaceAllString(s, "")
	fmt.Printf(processedString)
	i := 0
	j := len(processedString) - 1

	for i < j {
		if processedString[i] != processedString[j] {
			return false
		}
		i++
		j--
	}
	return true
}
