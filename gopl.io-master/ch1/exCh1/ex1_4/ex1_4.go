package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	duplicate := make(map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("No file names!")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			if countDuplicate(f) {
				duplicate[arg] = true
			} else {
				duplicate[arg] = false
			}
			f.Close()
		}
	}
	for file, hasDup:= range duplicate {
		if hasDup {
			fmt.Printf("%s\n", file)
		}
	}
}

func countDuplicate(f *os.File) bool  {
	countMap := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		countMap[input.Text()]++
		if countMap[input.Text()] > 1 {
			return true
		}
	}	
	return false
}
