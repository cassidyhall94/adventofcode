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

func addSums(scores []int) int {
	total := 0

	for i := 0; i < len(scores); i++ {
		total += scores[i]
	}
	return total
}

func assignPriorities(commonLetter string) []int {
	assignedPriorities := []int{}
	for _, letter := range commonLetter {
		if unicode.IsUpper(letter) {
			letter -= 38
			assignedPriorities = append(assignedPriorities, int(letter))
		}
		if unicode.IsLower(letter) {
			letter -= 96
			assignedPriorities = append(assignedPriorities, int(letter))
		}
	}
	return assignedPriorities
}
