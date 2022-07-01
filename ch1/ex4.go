package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
				continue
			}
			countLines(f, counts, filenames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t", n, line)
			for file := range filenames[line] {
				fmt.Printf("%s, ", file)
			}
			fmt.Println()
		}
	}
}
func countLines(f *os.File, counts map[string]int, filenames map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		filename := f.Name()
		// initialize the map for this line
		if _, ok := filenames[line]; !ok {
			filenames[line] = make(map[string]int)
		}
		filenames[line][filename]++
	}
}
