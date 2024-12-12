package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	CONSTANT_INCREASE = 1
	CONSTANT_DECREASE = 0
)

type Day2 struct{}

func (d *Day2) part1(scans *bufio.Scanner) {
	safe := 0
	for scans.Scan() {
		line := scans.Text()
		levelArr := strings.Split(line, " ")
		levels := make([]int, len(levelArr))
		for i, v := range levelArr {
			val, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(line, i, v)
				log.Fatalln(err)
			}
			levels[i] = val
		}
		isInc := -1
		hold := levels[0]
		isSafe := true
		for _, v := range levels[1:] {

			if isInc == -1 {
				if hold < v {
					isInc = 1
				} else {
					isInc = 0
				}
			}

			if isInc == CONSTANT_DECREASE && hold < v || isInc == CONSTANT_INCREASE && hold > v || absInt(hold-v) > 3 || hold == v {
				isSafe = false
				fmt.Println(levels)
				break
			}
			hold = v
		}
		if isSafe {
			safe++
		}

	}
	fmt.Println(safe)
	if err := scans.Err(); err != nil {
		log.Fatalln(scans.Err())
	}
}

func (d *Day2) dampUnsafe(levels []int, strikes int, isInc int) bool {
	levelsCopy := make([]int, len(levels))
	copy(levelsCopy, levels)

	hold := levelsCopy[0]
	mark := 0
	fault := 0
	for i, v := range levelsCopy[1:] {
		if isInc == -1 {
			if hold < v {
				isInc = CONSTANT_INCREASE
			} else {
				isInc = CONSTANT_DECREASE
			}
		}

		if isInc == CONSTANT_DECREASE && hold < v || isInc == CONSTANT_INCREASE && hold > v || absInt(hold-v) > 3 || hold == v {
			mark = i + 1
			strikes += 1
			fault += 1
		}
		hold = v
	}
	if strikes > 1 {
		return false
	}
	if strikes > 0 {
		if fault == 0 {
			return true
		}

		firstArr := make([]int, len(levelsCopy))
		secondArr := make([]int, len(levelsCopy))
		copy(firstArr, levelsCopy)
		copy(secondArr, levelsCopy)

		firstArr = append(firstArr[:mark], firstArr[mark+1:]...)
		secondArr = append(secondArr[:mark-1], secondArr[mark:]...)

		first := d.dampUnsafe(firstArr, strikes, isInc)
		second := d.dampUnsafe(secondArr, strikes, isInc)

		if !first && !second {
			return false
		}
	}
	return true
}

func (d *Day2) ScanUnsafe(levels []int) bool {
	hold := 0
	isInc := -1
	for i, v := range levels {
		if i == 0 {
			hold = v
		}
		if isInc == -1 && i > 0 {
			if hold < v {
				isInc = 1
			} else {
				isInc = 0
			}
		}

		if absInt(hold-v) > 3 || (isInc == CONSTANT_DECREASE && hold < v) || (isInc == CONSTANT_INCREASE && hold > v) || (hold == v && i != 0) {
			return false
		}
		hold = v
	}
	return true
}

func (d *Day2) part2(scans *bufio.Scanner) {
	safe := 0
	for scans.Scan() {
		line := scans.Text()
		levelArr := strings.Split(line, " ")
		levels := make([]int, len(levelArr))
		for i, v := range levelArr {
			val, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(line, i, v)
				log.Fatalln(err)
			}
			levels[i] = val
		}
		UndampendStatus := d.ScanUnsafe(levels)
		DampenedStatus := false
		if !UndampendStatus {
			for i := range levels {
				newArr := append([]int{}, levels[:i]...)
				newArr = append(newArr, levels[i+1:]...)
				DampenedStatus = d.ScanUnsafe(newArr)
				if DampenedStatus {
					break
				}
			}
		}
		if UndampendStatus || DampenedStatus {
			safe += 1
		}

	}
	fmt.Println(safe)
	if err := scans.Err(); err != nil {
		log.Fatalln(scans.Err())
	}
}
