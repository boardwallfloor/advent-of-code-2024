package main

import (
	"bufio"
	"fmt"
	"slices"
	"strings"
)

type Day5 struct {
	rule []string
}

func (d *Day5) addRule(input []string) {
	if len(d.rule) == 0 {
		d.rule = append(d.rule, input...)
		return
	}
	for i, v := range input {
		ok := slices.Index(d.rule, v)
		if ok == -1 {
			if i == 0 {
				d.rule = append([]string{v}, d.rule...)
			}
			if i == 1 {
				d.rule = append(d.rule, v)
			}
		}
	}
}

func (d *Day5) part1(scan *bufio.Scanner) {
	mode := 0
	for scan.Scan() {
		lines := scan.Text()
		if lines == "" {
			fmt.Println("-------")
			mode = 1
		} else if mode == 0 {
			input := strings.Split(lines, "|")
			d.addRule(input)
		} else {
		}
	}
	for _, v := range d.rule {
		fmt.Println(v)
	}
}
