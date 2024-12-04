package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	memory_list := ReadInput("./input-day3-mull_it_over.txt")
	fmt.Printf("Sum of multiplications in memory: %v\n", GetUncorruptedMul(memory_list, false))
	fmt.Printf("Sum of multiplications in memory with conditionals enabled: %v\n", GetUncorruptedMul(memory_list, true))
}

func GetUncorruptedMul(memory_list []string, conditional_enabled bool) int {
	regexp_instructions := `mul\((\d{1,3}),(\d{1,3})\)`
	if conditional_enabled {
		regexp_instructions += `|do\(\)|don't\(\)`
	}

	sum_of_mul := 0
	mul_instruction_enabled := true

	for _, memory_block := range memory_list {
		re := regexp.MustCompile(regexp_instructions)
		memory_instructions := re.FindAllStringSubmatch(memory_block, -1)

		for _, instruction := range memory_instructions {
			switch instruction[0] {
			case "do()":
				mul_instruction_enabled = true
			case "don't()":
				mul_instruction_enabled = false
			default:
				if mul_instruction_enabled {
					first_mul_num, _ := strconv.Atoi(instruction[1])
					second_mul_num, _ := strconv.Atoi(instruction[2])
					sum_of_mul += first_mul_num * second_mul_num
				}
			}
		}
	}

	return sum_of_mul
}

func ReadInput(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	input_list := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		input_list = append(input_list, line)
	}

	return input_list
}
