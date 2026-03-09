package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings" // 1. Added for string manipulation
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	file1, err := os.Open("example.txt")
	check(err)
	defer file1.Close()

	file2, err := os.Create("new.txt")
	check(err)
	defer file2.Close()

	writer := bufio.NewWriter(file2)
	file3 := bufio.NewScanner(file1)

	for file3.Scan() {
		line := file3.Text()

		// 2. Split the line into individual words
		words := strings.Fields(line)

		var modifiedWords []string
		for _, word := range words {
			// 3. Perform your manipulations here
			// Examples:
			upperWord := strings.ToUpper(word)
			// binWord := fmt.Sprintf("%b", someInteger) // If converting numbers

			modifiedWords = append(modifiedWords, upperWord)
		}

		// 4. Join words back with a space and write the line
		// Use Fprintln to keep the original line-by-line format
		fmt.Fprintln(writer, strings.Join(modifiedWords, " "))
	}

	err = writer.Flush()
	check(err)

	if err := file3.Err(); err != nil {
		log.Fatal(err)
	}
}
