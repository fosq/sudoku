package main

import (
	"fmt"
	"os"
)

var field [9][9]int

func main() {
	if len(os.Args) == 10 {
		arguments := os.Args[1:]
		field = inputToRowsAndColumns(arguments)
		if backtrack(&field) {
			printSudoku(field)
		} else {
			fmt.Println("Error")
		}
	} else {
		fmt.Println("Error")
	}
}

// backtracking algorithm
func backtrack(field *[9][9]int) bool {
	// last check for empty cells after backtrack reruns at line 60
	if !checkEmptyCell(field) {
		return true
	}
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if field[row][col] == 0 {
				// start finding substitute for zeros
				for substitute := 9; substitute > 0; substitute-- {
					field[row][col] = substitute
					// checks if field isn't OK with new substitute, set cell back to 0 and loop next to try to find next
					// integer that fits
					//
					// checks if field is OK, if that's positive, then reruns backtrack to check
					// if there aren't any empty cells, if there aren't any zeros, then returns true and prints field
					if fieldValidity(field) {
						if backtrack(field) {
							return true
						}
						field[row][col] = 0
					} else {
						field[row][col] = 0
					}
				}
				// if didn't find substitute, program fails
				return false
			}
		}
	}
	return false
}

// check if cell is empty
func checkEmptyCell(field *[9][9]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if field[row][col] == 0 {
				return true
			}
		}
	}
	return false
}

func fieldValidity(field *[9][9]int) bool {
	// check for duplicates in a row
	for row := 0; row < 9; row++ {
		// create 10 slot integer array for number counts in row/column/3x3
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			// count all the numbers in a row by columns
			counter[field[row][col]]++
		}
		// if duplicates found, return false
		if duplicateCheck(counter) {
			return false
		}
	}

	// check for duplicates in a column
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			// count all the numbers in a column by row
			counter[field[col][row]]++
		}
		if duplicateCheck(counter) {
			return false
		}
	}

	// check for duplicates in all 3x3s
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			// check individual 3x3s
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[field[row][col]]++
				}
				if duplicateCheck(counter) {
					return false
				}
			}
		}
	}

	return true
}

// check for duplicates in row/column/3x3 int counter
func duplicateCheck(counter [10]int) bool {
	for col := 1; col <= 9; col++ {
		if counter[col] > 1 {
			return true
		}
	}
	return false
}

// convert string array to bytes to int and return a new [9][9]int array
func inputToRowsAndColumns(array []string) [9][9]int {
	var newArray [9][9]int
	for x, row := range array {
		for z, column := range row {
			// . to 0
			if column == '.' {
				newArray[x][z] = 0
			} else {
				// byte to int
				newArray[x][z] = Atoi(string(column))
			}
		}
	}
	return newArray
}

// print le sudok
func printSudoku(array [9][9]int) {
	for i := range array {
		for j := range array[i] {
			if j != 8 {
				fmt.Print(array[i][j], " ")
			} else {
				fmt.Print(array[i][j])
			}
		}
		fmt.Println()
	}
}

// atoi (:
func Atoi(s string) int {
	runes := []rune(s)
	multi := 0
	isNeg := 0

	if s == "" {
		return 0
	}

	if runes[0] == '-' {
		isNeg = 1
	}

	if runes[0] == '-' || runes[0] == '+' {
		runes = runes[1:]
	}

	for i := 0; i < len(runes); i++ {
		if runes[i] < '0' || runes[i] > '9' {
			return 0
		}

		for j, k := 0, 0; j < 1; k++ {
			if runes[k] == '0' {
				if k == len(runes)-1 {
					return 0
				}
			} else {
				j++
			}
		}

		if runes[i] == '0' {
			multi = multi * 10
		} else {
			multi = multi*10 + int(runes[i]) - '0'
		}
	}
	if isNeg == 1 {
		multi = multi * -1
	}
	return multi
}
