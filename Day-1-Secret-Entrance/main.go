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

func Part1(input []string, startDial int) {
	var dial int = startDial
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

func Part2(input []string, startDial int) {
	var dial int = startDial
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

        step := 1
        if direction == "L" {
            step = -1
        }

        // Simulate step-by-step movement
        for i := 0; i < rotations; i++ {
            dial += step
            if dial < 0 {
                dial = 99
            } else if dial > 99 {
                dial = 0
            }

            // Count every time the dial points at 0
            if dial == 0 {
                count++
            }
        }
		fmt.Println(dial)
		fmt.Println()
    }
	fmt.Println(count)
}

func main() {
	input := ReadFile("input.txt")
	// input := []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}
	start := 50
	Part1(input, start)
	Part2(input, start)
}