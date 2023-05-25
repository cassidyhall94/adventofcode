package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

//Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?
//parse output file
// sum the grouped numbers together, seperate totals on a newline
// range over numbers to find the highest and return

func main() {
	file, err := os.Open("output.txt")
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
	addSums(ret)
}

func addSums(ret [][]int) int {
	totals := []int{}
	total := 0

	for _, elfSnacks := range ret {
		for i := 0; i < len(elfSnacks); i++ {
			total += elfSnacks[i]
		}
		totals = append(totals, total)
		total = 0
	}
	sort.Sort(sort.Reverse(sort.IntSlice(totals)))
	fmt.Println(totals[0])
	return totals[0]
}
