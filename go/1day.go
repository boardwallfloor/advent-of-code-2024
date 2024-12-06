package main

import (
	"bufio"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

type Day1 struct{}

func (d *Day1) part1(scan *bufio.Scanner) {
	list1 := make([]int, 0)
	list2 := make([]int, 0)
	for scan.Scan() {
		line := scan.Text()
		input := strings.Split(line, "   ")
		for i, v := range input {
			val, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(line, i, v)
				log.Fatalln(err)
			}
			if i == 0 {
				list1 = append(list1, val)
			}
			if i == 1 {
				list2 = append(list2, val)
			}
		}
	}
	slices.Sort(list1)
	slices.Sort(list2)
	sum := 0
	for i := range list1 {
		diff := list1[i] - list2[i]
		if diff < 0 {
			diff *= -1
		}
		sum += diff
	}
	fmt.Println(sum)
}

func (d *Day1) part2(scan *bufio.Scanner) {
	list1 := make([]int, 0)
	list2 := make([]int, 0)
	for scan.Scan() {
		line := scan.Text()
		input := strings.Split(line, "   ")
		for i, v := range input {
			val, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(line, i, v)
				log.Fatalln(err)
			}
			if i == 0 {
				list1 = append(list1, val)
			}
			if i == 1 {
				list2 = append(list2, val)
			}
		}
	}
	index := make(map[int]int, 0)
	for _, v := range list2 {
		if _, ok := index[v]; ok {
			index[v] += 1
		} else {
			index[v] = 1
		}
	}

	sum := 0
	for _, v := range list1 {
		total := index[v] * v
		sum += total
	}
	fmt.Println(sum)
}
