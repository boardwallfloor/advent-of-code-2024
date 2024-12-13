package main

import (
	"bufio"
	"fmt"
)

type Day4 struct {
	height     int
	width      int
	input      []string
	textLength int
	textSearch string
}

func (d *Day4) checkCardinal(x, y int) {
	directions := [][2]int{
		{-1, 0},  // North
		{-1, 1},  // Northeast
		{0, 1},   // East
		{1, 1},   // Southeast
		{1, 0},   // South
		{1, -1},  // Southwest
		{0, -1},  // West
		{-1, -1}, // Northwest
	}
	for _, v := range directions {
		height := y
		width := x
		count := 0
		for height < d.height || width < d.width || height > 0 || width > 0 && count < 5 {
			fmt.Print(height, width)
			val := d.input[height][width]
			if val != d.textSearch[count] {
				break
			}
			height += v[0]
			width += v[1]
			count++
		}
	}
}

func (d *Day4) part1(scan *bufio.Scanner) {
	textBox := make([]string, 0)
	for scan.Scan() {
		lines := scan.Text()
		textBox = append(textBox, lines)
	}
	d.width = len(textBox)
	d.height = len(textBox[0])
	d.input = textBox
	d.textLength = 5
	d.textSearch = "XMAS"
	fmt.Println(textBox)
	for y, line := range textBox {
		for x := range line {
			d.checkCardinal(x, y)
		}
	}
}
