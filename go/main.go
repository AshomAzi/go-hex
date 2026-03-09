package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	// 1. Wrap the output file in a buffered writer for efficiency
	writer := bufio.NewWriter(file2)

	file3 := bufio.NewScanner(file1)
	for file3.Scan() {
		line := file3.Text()
		
		// 2. We write to the buffer ('writer') instead of the raw file
		fmt.Fprintln(writer, line)
	}

	// 3. IMPORTANT: You MUST flush the buffer to ensure the last 
	// bits of data are actually written to the disk before the program ends.
	err = writer.Flush()
	check(err)

	if err := file3.Err(); err != nil {
		log.Fatal(err)
	}
}