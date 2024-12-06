package main

import (
	"bufio"
	"log"
	"os"
)

type ScannedInput struct {
	file *os.File
	scan *bufio.Scanner
}

func scanInput(filepath string) ScannedInput {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)
	return ScannedInput{file: file, scan: scanner}
}

func absInt(result int) int {
	if result < 0 {
		return -result
	}
	return result
}
