package main

import (
	"bufio"
	"os"
)

func main() {

	// opening the file in read-only mode. The file must exist (in the current working directory)
	file, _ := os.Open("chars.txt")

	// the file value returned by os.Open() is wrapped in a bufio.Scanner just like a buffered reader.
	scanned := bufio.NewScanner(file)

	scanned.Split(bufio.ScanLines)

	var lines []string

	for scanned.Scan() {
		lines = append(lines, scanned.Text())
	}

	file.Close()
}
