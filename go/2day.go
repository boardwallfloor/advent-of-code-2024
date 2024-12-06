package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
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

			if isInc == 0 && hold < v || isInc == 1 && hold > v || absInt(hold-v) > 3 || hold == v {
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
	// Base case: A one- or two-element array is always safe
	if len(levels) <= 2 {
		return true
	}

	// Check if the current levels array is safe
	hold := levels[0]
	for i := 1; i < len(levels); i++ {
		diff := absInt(levels[i] - hold)

		// Determine if the sequence is valid
		if diff < 1 || diff > 3 || levels[i] == hold ||
			(isInc == 0 && levels[i] > hold) ||
			(isInc == 1 && levels[i] < hold) {
			// Sequence is invalid; decide if a strike is tolerable
			strikes++
			if strikes > 1 {
				return false // More than one strike makes it unsafe
			}

			// Try removing any one level and check for safety
			for j := 0; j < len(levels); j++ {
				modifiedLevels := append(levels[:j], levels[j+1:]...)
				if d.dampUnsafe(modifiedLevels, strikes, -1) {
					return true
				}
			}
			return false // No valid removal found
		}

		// Update direction (isInc) if not yet determined
		if isInc == -1 {
			if levels[i] > hold {
				isInc = 1
			} else if levels[i] < hold {
				isInc = 0
			}
		}
		hold = levels[i]
	}

	return true // No faults found, sequence is safe
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
		isSafe := d.dampUnsafe(levels, 0, -1)
		if isSafe {
			safe++
		}

	}
	fmt.Println(safe)
	if err := scans.Err(); err != nil {
		log.Fatalln(scans.Err())
	}
}

// package main
//
// import (
// 	"fmt"
// 	"math"
// )
//
// // Helper function to check if a slice is monotonic and has valid level differences
// func isMonotonicAndSafe(levels []int) bool {
// 	if len(levels) < 2 {
// 		return true // A single level is trivially safe
// 	}
//
// 	isIncreasing := true
// 	isDecreasing := true
//
// 	for i := 0; i < len(levels)-1; i++ {
// 		diff := levels[i+1] - levels[i]
// 		if diff < 1 || diff > 3 {
// 			return false // Invalid level difference
// 		}
// 		if levels[i+1] < levels[i] {
// 			isIncreasing = false
// 		}
// 		if levels[i+1] > levels[i] {
// 			isDecreasing = false
// 		}
// 	}
//
// 	return isIncreasing || isDecreasing
// }
//
// // Function to check if removing one level can make the report safe
// func isSafeWithOneRemoval(levels []int) bool {
// 	for i := 0; i < len(levels); i++ {
// 		// Create a new slice excluding the current level
// 		modifiedLevels := append(levels[:i], levels[i+1:]...)
// 		if isMonotonicAndSafe(modifiedLevels) {
// 			return true
// 		}
// 	}
// 	return false
// }
//
// // Main function to analyze reports
// func analyzeReport(levels []int) string {
// 	if isMonotonicAndSafe(levels) || isSafeWithOneRemoval(levels) {
// 		return "Safe"
// 	}
// 	return "Unsafe"
// }
//
// func main() {
// 	// Test cases
// 	reports := [][]int{
// 		{7, 6, 4, 2, 1},   // Safe
// 		{1, 2, 4, 7, 3},   // Unsafe
// 		{1, 3, 6, 5, 2},   // Unsafe
// 		{10, 8, 6, 4, 3},  // Safe
// 		{5, 6, 9, 12},     // Unsafe
// 		{3, 6, 9, 12, 15}, // Safe
// 	}
//
// 	for _, report := range reports {
// 		fmt.Printf("Report %v: %s\n", report, analyzeReport(report))
// 	}
// }
