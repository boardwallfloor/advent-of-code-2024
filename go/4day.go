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

func (d *Day4) checkCardinal(x, y int) int {
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
	total := 0
	for _, v := range directions {
		height := y
		width := x
		count := 0
		temp := ""
		for count < 4 {
			if height >= d.height || width >= d.width || height < 0 || width < 0 {
				break
			}
			val := d.input[height][width]
			if val != d.textSearch[count] {
				break
			}
			height += v[0]
			width += v[1]
			temp += string(d.textSearch[count])
			count++
		}
		if count == 4 {
			total += 1
			// fmt.Println(v, y, x, height, width, temp)
		}
	}
	return total
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
	total := 0
	for y, line := range textBox {
		for x := range line {
			total += d.checkCardinal(x, y)
		}
	}
	fmt.Println(total)
}

func (d *Day4) checkCross(x, y int) int {
	if d.input[y][x] != 'A' {
		return 0
	}
	directions := [2][4]int{
		{-1, -1, 1, 1}, // Top-left // Bottom-right
		{-1, 1, 1, -1}, // Top-right // Bottom-left
	}
	tot := 0
	temp := [3][3]string{}
	temp[1][1] = "A"
	for _, v := range directions {

		height := []int{
			y + v[0],
			y + v[2],
		}
		width := []int{
			x + v[1],
			x + v[3],
		}
		if height[0] >= d.height || width[0] >= d.width || height[0] < 0 || width[0] < 0 {
			if y == 7 {
				// fmt.Print(".")
				// fmt.Println("OUCH")
			}
			return 0
		}
		if height[1] >= d.height || width[1] >= d.width || height[1] < 0 || width[1] < 0 {
			if y == 7 {
				// fmt.Print(".")
				// fmt.Printf("Condition triggered: height[1]=%d, d.height=%d, width[1]=%d, d.width=%d\n", height[1], d.height, width[1], d.width)
				// fmt.Println("OUWIE")
			}
			return 0
		}
		if d.input[height[0]][width[0]] == 'M' && d.input[height[1]][width[1]] == 'S' {
			temp[1+v[0]][1+v[1]] = "M"
			temp[1+v[2]][1+v[3]] = "S"
			tot += 1
		} else if d.input[height[0]][width[0]] == 'S' && d.input[height[1]][width[1]] == 'M' {
			temp[1+v[0]][1+v[1]] = "S"
			temp[1+v[2]][1+v[3]] = "M"
			tot += 1
		}
	}
	// if y == 7 {
	// 	for _, v := range temp {
	// 		fmt.Println(v)
	// 	}
	// 	fmt.Println("==")
	// }
	if tot == 2 {
		return 1
	}
	return 0
}

func (d *Day4) part2(scan *bufio.Scanner) {
	textBox := make([]string, 0)
	for scan.Scan() {
		lines := scan.Text()
		textBox = append(textBox, lines)
	}
	d.width = len(textBox)
	d.height = len(textBox[0])
	d.input = textBox
	total := 0
	// a := 0
	for y, line := range textBox {
		// if y == 7 {
		// 	fmt.Println(line)
		// }
		for x := range line {
			// if y == 7 {
			// 	if d.input[y][x] == 'A' {
			// 		// a++
			// 		fmt.Println("A")
			// 	} else {
			// 		fmt.Println(".")
			// 	}
			// }
			stuff := d.checkCross(x, y)
			total += stuff
			// if stuff == 1 {
			// 	fmt.Print("A")
			// } else {
			// 	fmt.Print(".")
			// }
		}
	}
	// fmt.Printf("A = %d\n", a)
	fmt.Println(total)
}
