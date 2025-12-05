package main

import "fmt"
import "os"
import "strings"
import "strconv"

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

func main() {
	input := ReadFile("input.txt")
	// testInput := []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}
	var dial int = 50
	var count int = 0

	for _, instruction := range input {
		var direction string = string(instruction[0])
		fmt.Println(direction)
		var raw_rotations string = string(instruction[1:])
		rotations, err := strconv.Atoi(raw_rotations)
		if err != nil {
			fmt.Println("Error parsing rotations as int")
		}
		fmt.Println(rotations)
		fmt.Println()

		if direction == "R" {
			dial += rotations
		} else {
			dial -= rotations
		}

		for dial < 0 || dial > 99 {
			if dial < 0 {
				dial = 100 + dial
			} else if dial > 99 {
				dial = dial - 100
			}
		}

		if dial == 0 {
			count += 1
		}

		fmt.Println(dial)
		fmt.Println()
	}

	fmt.Println(count)
}