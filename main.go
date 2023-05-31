package main

import (
	"fmt"
	"sort"
)

func main() {
	// ret_day1 := readFile("day1_input.txt")
	ret_day2 := readFile2("day2_input.txt")
	// totalCalories := day1_addSums(ret_day1)
	totalScore := day2_addScores(ret_day2)
	totalNewScore := day2_addNewScores(ret_day2)
	// fmt.Println(totalCalories)
	fmt.Println(totalScore)
	fmt.Println(totalNewScore)
}

func day1_addSums(ret_day1 [][]int) int {
	totals := []int{}
	total := 0

	for _, elfSnacks := range ret_day1 {
		for i := 0; i < len(elfSnacks); i++ {
			total += elfSnacks[i]
		}
		totals = append(totals, total)
		total = 0
	}
	sort.Sort(sort.Reverse(sort.IntSlice(totals)))
	// topThreeCalories := day1_addTopThree(totals)
	// fmt.Println(topThreeCalories)
	return totals[0]
}

func day1_addTopThree(summedCalories []int) int {
	total := 0

	for i, elfSnacks := range summedCalories {
		if i < 3 {
			total += elfSnacks
		}
	}
	return total
}

func day2_addScores(ret_day2 [][]string) int {
	total := 0
	paper := 2
	rock := 1
	scissors := 3

	lose := 0
	draw := 3
	win := 6

	for _, round := range ret_day2 {
		for _, turn := range round {
			for _, char := range turn {
				if char != 32 {
					if turn[0] == 65 {
						if string(char) == "X" {
							total += rock + draw
						}
						if string(char) == "Y" {
							total += paper + win
						}
						if string(char) == "Z" {
							total += scissors + lose
						}
					}
					if turn[0] == 66 {
						if string(char) == "X" {
							total += rock + lose
						}
						if string(char) == "Y" {
							total += paper + draw
						}
						if string(char) == "Z" {
							total += scissors + win
						}
					}
					if turn[0] == 67 {
						if string(char) == "X" {
							total += rock + win
						}
						if string(char) == "Y" {
							total += paper + lose
						}
						if string(char) == "Z" {
							total += scissors + draw
						}
					}
				}
			}
		}
	}
	return total
}

func day2_addNewScores(ret_day2 [][]string) int {
	total := 0
	paper := 2
	rock := 1
	scissors := 3

	lose := 0
	draw := 3
	win := 6

	for _, round := range ret_day2 {
		for _, turn := range round {
			for _, char := range turn {
				if char != 32 {
					if turn[0] == 65 {
						if string(char) == "X" {
							total += scissors + lose
						}
						if string(char) == "Y" {
							total += rock + draw
						}
						if string(char) == "Z" {
							total += paper + win
						}
					}
					if turn[0] == 66 {
						if string(char) == "X" {
							total += rock + lose
						}
						if string(char) == "Y" {
							total += paper + draw
						}
						if string(char) == "Z" {
							total += scissors + win
						}
					}
					if turn[0] == 67 {
						if string(char) == "X" {
							total += paper + lose
						}
						if string(char) == "Y" {
							total += scissors + draw
						}
						if string(char) == "Z" {
							total += rock + win
						}
					}
				}
			}
		}
	}
	return total
}
