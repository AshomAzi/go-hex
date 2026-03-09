package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"newgo/conv" // Use your module name here
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
	scanner := bufio.NewScanner(file1)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		var modifiedWords []string

		for i := 0; i < len(words); i++ {
			// If this word contains "(hex)" and there is a word before it
			if i > 0 && strings.Contains(words[i], "(hex)") {
				// We replace the PREVIOUS word in our result slice
				// with the converted value
				lastIdx := len(modifiedWords) - 1
				modifiedWords[lastIdx] = conv.ProcessHex(words[i-1], words[i])
				
				// Optional: If you want to delete the "(hex)" word itself, 
				// you can skip adding it to the slice.
				// If you want to keep it, uncomment the line below:
				// modifiedWords = append(modifiedWords, words[i])
			} else {
				modifiedWords = append(modifiedWords, words[i])
			}
		}

		fmt.Fprintln(writer, strings.Join(modifiedWords, " "))
	}

	writer.Flush()
}