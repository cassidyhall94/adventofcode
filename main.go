package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// ret_day1 := readFile("day1_input.txt")
	// ret_day2 := readFile2("day2_input.txt")
	// ret_day3 := readFile2("day3_input.txt")
	// ret_day4 := readFile3("day4_input.txt")
	ret_day5 := readFile3("day5_input.txt")

	// totalCalories := day1_addSums(ret_day1)

	// totalScore := day2_addScores(ret_day2)
	// totalNewScore := day2_addNewScores(ret_day2)

	// totalPriorities := day3_sumPriorities(ret_day3)
	// totalThreeElves := day3_sumElvesPriorities(ret_day3)

	// calculatedPairs := day4_calculatePairs(ret_day4)
	// overlappingPairs := day4_calculateOverlaps(ret_day4)

	sortedStacks := day5_sortStackNums(ret_day5)
	sortedResult := day5_crateStackArrays(sortedStacks)
	fmt.Println(sortedResult)
	// crate := day5_crates(ret_day5)

	// fmt.Println(totalCalories)
	// fmt.Println(totalScore)
	// fmt.Println(totalNewScore)
	// fmt.Println(totalPriorities)
	// fmt.Println(totalThreeElves)
	// fmt.Println(calculatedPairs)
	// fmt.Println(overlappingPairs)
	// fmt.Println(crate)
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

func day3_sumPriorities(ret_day3 [][]string) int {
	count := 0
	total := 0
	commonLetter := ""
	for _, rucksacks := range ret_day3 {
		for _, rucksack := range rucksacks {
			halfLen := len(rucksack) / 2
			item1 := rucksack[:halfLen]
			item2 := rucksack[halfLen:]
			for _, item1Char := range item1 {
				for _, item2Char := range item2 {
					if item1Char == item2Char && count == 0 {
						count++
						commonLetter += string(item2Char)
					}
				}
			}
			count = 0
		}
	}
	priorities := assignPriorities(commonLetter)
	total = addSums(priorities)
	return total
}

func day3_sumElvesPriorities(ret_day3 [][]string) int {
	group := []string{}
	groups := [][]string{}
	for _, rucksacks := range ret_day3 {
		for _, rucksack := range rucksacks {
			fmt.Println("RUCKSACK:", rucksack)
			if len(group) < 3 {
				group = append(group, rucksack)
				fmt.Println("GROUP:", group)
				// fmt.Println("GROUPS: ", len(group))
			}
			if len(group) == 3 {
				groups = append(groups, group)
				group = []string{}
			}
		}

	}
	fmt.Println("GROUPS", groups)
	commonLetters := day3_compareGroups(groups)
	priorities := assignPriorities(commonLetters)
	total := addSums(priorities)
	return total
}

func day3_compareGroups(groups [][]string) string {
	commonLetters := ""
	for _, elfGroup := range groups {
		elfGroup := elfGroup
		firstElf := elfGroup[0]
		secondElf := elfGroup[1]
		thirdElf := elfGroup[2]
		for _, elf1Char := range firstElf {
			if strings.Contains(secondElf, string(elf1Char)) && strings.Contains(thirdElf, string(elf1Char)) {
				commonLetters += string(elf1Char)
				break
			}
		}
	}
	return commonLetters
}

func day4_calculatePairs(ret_day4 []string) int {
	count := 0
	for _, pairs := range ret_day4 {
		pair := strings.Split(pairs, ",")
		assignment1 := fillNumbers(pair[0])
		assignment2 := fillNumbers(pair[1])
		if containsAssignment(assignment1, assignment2) || containsAssignment(assignment2, assignment1) {
			count++
		}
	}
	return count
}

func containsAssignment(assignment1, assignment2 []string) bool {
	for _, num := range assignment2 {
		found := false
		for _, n := range assignment1 {
			if num == n {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func fillNumbers(assignment string) []string {
	numbers := strings.Split(assignment, "-")
	start, _ := strconv.Atoi(numbers[0])
	end, _ := strconv.Atoi(numbers[1])

	result := []string{}
	for i := start; i <= end; i++ {
		result = append(result, strconv.Itoa(i))
	}
	return result
}

func day4_calculateOverlaps(ret_day4 []string) int {
	count := 0
	for _, pairs := range ret_day4 {
		pair := strings.Split(pairs, ",")
		assignment1 := fillNumbers(pair[0])
		assignment2 := fillNumbers(pair[1])
		if containsOverlap(assignment1, assignment2) || containsOverlap(assignment2, assignment1) {
			count++
		}
	}
	return count
}

func containsOverlap(assignment1, assignment2 []string) bool {
	for _, num := range assignment1 {
		fmt.Println(num)
		for _, n := range assignment2 {
			fmt.Println(n)
			if num == n {
				return true
			}

		}
	}
	return false
}

func day5_crates(ret_day5 []string) string {
	// for i, line := range ret_day5 {
	// 	for _, ln := range line {
	// 		fmt.Println(i, ln)
	// 	}
	// }
	return ""
}

func day5_sortStackNums(ret_day5 []string) []string {
	result := []string{}
	for _, line := range ret_day5 {
		for j, byte := range line {
			if string(byte) == " " && j == 0 {
				result = append(result, line)
			}
		}
	}
	return result
}

func day5_crateStackArrays(stackNums []string) [][]string {
	result := [][]string{}
	for _, nums := range stackNums {
		fmt.Println(nums)
		splitNums := strings.Split(nums, " ")
		fmt.Printf("%q\n", splitNums)

	}
	return result
}
