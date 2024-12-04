package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	memory_list := ReadInput("./input-day3-mull_it_over.txt")
	fmt.Printf("Sum of multiplications in memory: %v\n", GetUncorruptedMul(memory_list, false))
	fmt.Printf("Sum of multiplications in memory with conditionals enabled: %v\n", GetUncorruptedMul(memory_list, true))
}

func GetUncorruptedMul(memory_list []string, conditional_enabled bool) int {
	sum_of_mul := 0
	mul_instruction_enabled := true
	for _, memory_block := range memory_list {
		for idx := 0; idx < len(memory_block); idx++ {
			first_mul_num := ""
			second_mul_num := ""
			mul_is_valid := false

			if conditional_enabled {
				mul_instruction_enabled = FindConditional(idx, memory_block, mul_instruction_enabled)
			}

			if mul_instruction_enabled {
				first_mul_num, second_mul_num, mul_is_valid = FindMultiplication(idx, memory_block, first_mul_num, second_mul_num, mul_is_valid)
			}

			if mul_is_valid {
				first_mul_num, err := strconv.Atoi(first_mul_num)
				if err != nil {
					log.Fatal(err)
				}
				second_mul_num, err := strconv.Atoi(second_mul_num)
				if err != nil {
					log.Fatal(err)
				}

				sum_of_mul += first_mul_num * second_mul_num
			}
		}
	}
	return sum_of_mul
}

func FindConditional(idx int, memory_block string, mul_instruction_enabled bool) bool {
	if memory_block[idx] == 'd' {
		if memory_block[idx+1] == 'o' {
			switch memory_block[idx+2] {
			case '(':
				if memory_block[idx+3] == ')' {
					mul_instruction_enabled = true
				}
			case 'n':
				if memory_block[idx+3] == '\'' {
					if memory_block[idx+4] == 't' {
						if memory_block[idx+5] == '(' {
							if memory_block[idx+6] == ')' {
								mul_instruction_enabled = false
							}
						}
					}
				}
			}
		}
	}
	return mul_instruction_enabled
}

func FindMultiplication(idx int, memory_block string, first_mul_num string, second_mul_num string, mul_is_valid bool) (string, string, bool) {
	if memory_block[idx] == 'm' {
		if memory_block[idx+1] == 'u' {
			if memory_block[idx+2] == 'l' {
				if memory_block[idx+3] == '(' {
					for nidx := 1; nidx < 4; nidx++ {
						if memory_block[idx+3+nidx] >= '0' &&
							memory_block[idx+3+nidx] <= '9' {
							first_mul_num += string(memory_block[idx+3+nidx])
						}
						if memory_block[idx+3+nidx] == ',' &&
							nidx != 1 {
							break
						}
					}
					if first_mul_num != "" &&
						memory_block[idx+4+len(first_mul_num)] == ',' {
						for nidx := 1; nidx < 4; nidx++ {
							if memory_block[idx+4+len(first_mul_num)+nidx] >= '0' &&
								memory_block[idx+4+len(first_mul_num)+nidx] <= '9' {
								second_mul_num += string(memory_block[idx+4+len(first_mul_num)+nidx])
							}
							if memory_block[idx+4+len(first_mul_num)+nidx] == ')' &&
								nidx != 1 {
								break
							}
							if second_mul_num != "" &&
								memory_block[idx+5+len(first_mul_num)+len(second_mul_num)] == ')' {
								mul_is_valid = true
							}
						}
					}
				}
			}
		}
	}
	return first_mul_num, second_mul_num, mul_is_valid
}

func ReadInput(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	input_list := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		input_list = append(input_list, line)

	}

	defer file.Close()

	return input_list
}
