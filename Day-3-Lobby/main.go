package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func ReadFile(filename string) ([]string) {
	data, err := os.ReadFile(filename)
	dataAsStr := string(data)
	arr := strings.Split(dataAsStr, "\n")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	return arr
}

func getMax(numAsStr string) (string, int) {
	max := 0
	maxIndex := -1
	for i := 0; i < len(numAsStr); i++ {
		num, intConvertErr := strconv.Atoi(string(numAsStr[i]))
		if intConvertErr != nil {
			fmt.Println("formatting error")
		}
		if num > max {
			max = num
			maxIndex = i
		}
	}
	return strconv.Itoa(max), maxIndex
}

func main() {
	input := ReadFile("input.txt")
	// input := []string { "987654321111111", "811111111111119", "234234234234278", "818181911112111" }

	total := 0
	for _, bank := range input {
		max, maxIndex := getMax(bank)
		newBank := bank[maxIndex+1:]
		swap := false
		if maxIndex+1 == len(bank) {
			swap = true
			newBank = bank[:maxIndex]
		}
		secondMax, secondMaxIndex := getMax(newBank)
		if secondMaxIndex < 0 {
			fmt.Println("Error")
			break
		}
		largestJoltageAsStr := max + secondMax
		if swap {
			largestJoltageAsStr = secondMax + max
		}
		largestJoltage, convertErr := strconv.Atoi(largestJoltageAsStr)
		if convertErr != nil {
			fmt.Println("formatting error")
		}
		total += largestJoltage
	}
	fmt.Println(total)
}