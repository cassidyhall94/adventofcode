package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func readFile(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("file open error: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cur := []int{}
	ret := [][]int{}

	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			convStr, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("strconv error: %w", err)
			}
			cur = append(cur, convStr)
		}

		if len(scanner.Text()) == 0 {
			ret = append(ret, cur)
			cur = []int{}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("file scanner error: %w", err)
	}
	return ret
}

func readFile2(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("file open error: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cur := []string{}
	ret := [][]string{}

	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			cur = append(cur, scanner.Text())
		}
		ret = append(ret, cur)
		cur = []string{}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("file scanner error: %w", err)
	}
	return ret
}

func readFile3(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("file open error: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cur := ""
	ret := []string{}

	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			cur += scanner.Text()
		}
		ret = append(ret, cur)
		cur = ""
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("file scanner error: %w", err)
	}
	return ret
}

func addSums(scores []int) int {
	total := 0

	for i := 0; i < len(scores); i++ {
		total += scores[i]
	}
	return total
}

func assignPriorities(commonLetter string) []int {
	assignedPriorities := []int{}
	var assigned int
	for _, letter := range commonLetter {
		if unicode.IsUpper(letter) {
			letter -= 38
			assigned = int(letter)
			assignedPriorities = append(assignedPriorities, assigned)
		}
		if unicode.IsLower(letter) {
			letter -= 96
			assigned = int(letter)
			assignedPriorities = append(assignedPriorities, assigned)
			fmt.Println(assignedPriorities)
		}
	}
	return assignedPriorities
}

func reverseStrings(slice []string) []string {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func countLetters(crate string) int {
	count := 0
	for _, cra := range crate {
		if unicode.IsLetter(cra) {
			count++
		}
	}
	return count
}
