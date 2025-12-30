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

func Part2PatternCheck(input string) (bool) {
	inputLen := len(input)
	for i := 1; i < inputLen; i++ {
		match := true
		sequence := string(input[:i])
		sequenceLen := len(sequence)
		if inputLen % len(sequence) != 0 {
			continue
		}
		for j := sequenceLen; j < inputLen; j+=sequenceLen {
			finalSequenceIndex := j+sequenceLen
			checkSequence := string(input[j:finalSequenceIndex])
			if checkSequence != sequence {
				match = false
				break
			}
		}
		if match == true {
			// fmt.Println("match seq", sequence)
			return true
		}
	}
	return false
}

func Part1PatternCheck(input string) (bool) {
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
			isPattern := Part2PatternCheck(iAsStr)
			if isPattern {
				// fmt.Println("Invalid id found", i)
				// fmt.Println("=========")
				total += i
			}
		}
	}

	fmt.Println(total)
}