package main

/*
1.	Reading a single word using fmt.Scan

2.	Reading a single line using fmt.Scanln

3.	Reading a full sentence (with spaces) using bufio.NewReader

4.	Clean separation of functions with robust error handling

*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Reads a single word (stops at whitespace)
func readSingleWord() (string, error) {
	var word string
	fmt.Print("Enter a single word: ")
	_, err := fmt.Scan(&word)
	if err != nil {
		return "", fmt.Errorf("error reading word: %w", err)
	}
	return word, nil
}

// Reads words until a newline (space-separated, no newline support inside input)
func readWordsLine() (string, error) {
	var line string
	fmt.Print("Enter words (single line, space-separated): ")
	_, err := fmt.Scanln(&line)
	if err != nil {
		return "", fmt.Errorf("error reading line: %w", err)
	}
	return line, nil
}

// Reads the full line (with spaces), includes newline `\n` at the end
func readFullLine() (string, error) {
	fmt.Print("Enter a full sentence: ")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error reading full line: %w", err)
	}
	return strings.TrimSpace(line), nil // Remove trailing newline
}

func main() {
	// Read single word
	word, err := readSingleWord()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read word: %v\n", err)
	} else {
		fmt.Println("You entered word:", word)
	}

	// Clear the scanner buffer before next scan (useful after fmt.Scan)
	fmt.Scanln()

	// Read line with fmt.Scanln
	line, err := readWordsLine()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read line (Scanln): %v\n", err)
	} else {
		fmt.Println("You entered line (Scanln):", line)
	}

	// Read full sentence with bufio
	full, err := readFullLine()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read full line: %v\n", err)
	} else {
		fmt.Println("You entered sentence:", full)
	}
}
