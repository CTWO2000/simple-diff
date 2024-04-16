package main

import (
	"flag"
	"fmt"
	"slices"

	readfile "github.com/CTWO2000/simple-diff/utils/readFile"
)

func main() {

	primaryFilePath := flag.String("primary", "primary.txt", "Path to primary file")
	secondaryFilePath := flag.String("secondary", "secondary.txt", "Path to secondary file")

	flag.Parse()

	primaryFile := readfile.ReadFile(*primaryFilePath)
	secondaryFile := readfile.ReadFile(*secondaryFilePath)

	for _, line := range primaryFile {
		if !slices.Contains(secondaryFile, line) {
			fmt.Println(line)
		}
	}
}
