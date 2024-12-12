package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Day3 struct{}

func (d *Day3) isDigitOrComma(ch byte) bool {
	return '0' <= ch && ch <= '9' || ch == ','
}

func (d *Day3) parseAndOutput(input string) int {
	param := input[4 : len(input)-1]
	arr := strings.Split(param, ",")
	num1, err := strconv.Atoi(arr[0])
	if err != nil {
		log.Fatal(err)
	}
	num2, err := strconv.Atoi(arr[1])
	if err != nil {
		log.Fatal(err)
	}
	return num1 * num2
}

func (d *Day3) part2(scans *bufio.Scanner) {
	output := 0
	skip := false
	for scans.Scan() {
		codeString := scans.Text()
		for i, v := range codeString {
			if v == 'd' {
				if codeString[i:i+4] == "do()" {
					skip = false
				} else if codeString[i:i+7] == "don't()" {
					skip = true
				}
			}
			if v == 'm' && !skip {
				if codeString[i:i+4] != "mul(" {
					continue
				}

				for j := i + 4; j < len(codeString); j++ {
					if codeString[j] == ')' {
						validString := codeString[i : j+1]
						if !strings.ContainsRune(validString, ',') {
							break
						}
						res := d.parseAndOutput(validString)
						output += res
						break
					}
					if !d.isDigitOrComma(codeString[j]) {
						break
					}

				}
			}
		}
	}
	fmt.Println(output)
}

func (d *Day3) part1(scans *bufio.Scanner) {
	output := 0
	for scans.Scan() {
		codeString := scans.Text()
		for i, v := range codeString {
			if v == 'm' {
				if codeString[i:i+4] != "mul(" {
					continue
				}

				for j := i + 4; j < len(codeString); j++ {
					if codeString[j] == ')' {
						validString := codeString[i : j+1]
						if !strings.ContainsRune(validString, ',') {
							break
						}
						res := d.parseAndOutput(validString)
						output += res
						break
					}
					if !d.isDigitOrComma(codeString[j]) {
						break
					}

				}
			}
		}
	}
	fmt.Println(output)
}
