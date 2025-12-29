package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string) ([]string) {
	data, err := os.ReadFile(filename)
	dataAsStr := string(data)
	arr := strings.Split(dataAsStr, ",")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	return arr
}

func pattern_check(input string) (bool) {
	inputLen := len(input)
	if inputLen % 2 != 0 {
		return false
	}

	halfInputLen := inputLen/2
	firstHalf := string(input[0:halfInputLen])
	secondHalf := string(input[halfInputLen:])

	if firstHalf == secondHalf {
		return true
	}
	return false
}

func main() {
	input := ReadFile("input.txt")
	// input := []string{"11-22", "95-115", "998-1012" , "1188511880-1188511890", "222220-222224", "1698522-1698528", "446443-446449", "38593856-38593862", "565653-565659", "824824821-824824827", "2121212118-2121212124"}
	// fmt.Println(input)
	total := 0

	for _, id := range input {
		id_range := strings.Split(id, "-")
		start, startErr := strconv.Atoi(id_range[0])
		end, endErr := strconv.Atoi(id_range[1])
		if startErr != nil || endErr != nil {
			fmt.Println("formatting error")
		}

		for i := start; i <= end; i++ {
			iAsStr := strconv.Itoa(i)
			isPattern := pattern_check(iAsStr)
			if isPattern {
				total += i
			}
		}
	}

	fmt.Println(total)
}