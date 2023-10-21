// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	duplicateFileNames := make(map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, duplicateFileNames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, duplicateFileNames)
			f.Close()
		}
		for fileName, isDuplicate := range duplicateFileNames {
			if isDuplicate {
				fmt.Println("File with duplicates:", fileName)
			}
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, duplicateFileNames map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if !duplicateFileNames[f.Name()] && counts[input.Text()] > 0 {
			duplicateFileNames[f.Name()] = true
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
