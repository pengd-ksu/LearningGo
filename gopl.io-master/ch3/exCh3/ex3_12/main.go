package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please provide two arguments!")
		os.Exit(1)
	}
	if anagram(os.Args[1], os.Args[2]) {
		fmt.Printf("The two strings %s and %s are anagrams"+
			"of each other.\n", os.Args[1], os.Args[2])
	} else {
		fmt.Printf("The two strings %s and %s are not anagrams"+
			"of each other.\n", os.Args[1], os.Args[2])
	}
}

// The overralpping parts of comparing length of strings in the
// following functions, is to keep their capabilities of recyling
// under other senarios.
func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	// s1 and s2 must have the same length
	if sameStrings(s1, s2) {
		return false
	}
	return equalFrequency(collectFrequency(s1), collectFrequency(s2))
}

func sameStrings(s1, s2 string) bool {
	// Compare if two strings are equal
	if len(s1) != len(s2) {
		return false
	}
	var limit int = len(s1) - 1
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[limit-i] {
			return false
		}
	}
	return true
}

func equalFrequency(m1, m2 map[rune]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}
	return true
}

func collectFrequency(s string) map[rune]int {
	frequency := make(map[rune]int)
	// r will be a rune, can't use s[index]
	for _, r := range s {
		frequency[r]++
	}
	return frequency
}
