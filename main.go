package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// opening the file in read-only mode. The file must exist (in the current working directory)
	file, _ := os.Open("standard.txt")

	// the file value returned by os.Open() is wrapped in a bufio.Scanner just like a buffered reader.
	scanned := bufio.NewScanner(file)

	scanned.Split(bufio.ScanLines)

	var lines []string

	for scanned.Scan() {
		lines = append(lines, scanned.Text())
	}

	file.Close()

	asciiChrs := make(map[int][]string)

	dec := 31

	for _, line := range lines {
		if line == "" {
			dec++
		} else {
			asciiChrs[dec] = append(asciiChrs[dec], line)
		}
	}

	args := os.Args[1]

	for i := 0; i < len(args); i++ {
		if args[i] == 92 && args[i+1] == 110 {
			Newline(string(args[:i]), asciiChrs)
			Newline(string(args[i+2:]), asciiChrs)

		}
	}

	if strings.Contains(args, "\\n") == false {
		Newline(args, asciiChrs)
	}

}

func Newline(n string, y map[int][]string) {

	//prints horizontally

	for j := 0; j < len(y[32]); j++ {
		for _, letter := range n {
			fmt.Print(y[int(letter)][j])
		}
		fmt.Println()
	}
}
