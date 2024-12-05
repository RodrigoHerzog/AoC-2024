package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := ReadInput("./input-day5-print_queue.txt")
	rules, updates := FormatInput(input)
	correct_updates, incorrect_updates := VerifyUpdates(updates, rules)
	fixed_updates := FixUpdates(incorrect_updates, rules)

	fmt.Printf("Sum of correct updates middle number: %v\n", SumMiddle(correct_updates))
	fmt.Printf("Sum of fixed incorrect updates middle number: %v\n", SumMiddle(fixed_updates))
}

func FixUpdates(incorrect_updates [][]int, rules map[int][]int) [][]int {
	fixed_updates := [][]int{}
	for _, incorrect_update := range incorrect_updates {
		fixed_update := []int{}
		for _, incorrect_update_num := range incorrect_update {
			inserted_at := -1
			for _, rule := range rules[incorrect_update_num] {
				for fixed_update_idx, fixed_update_num := range fixed_update {
					if rule == fixed_update_num {
						if inserted_at == -1 {
							fixed_update = slices.Insert(fixed_update, fixed_update_idx, incorrect_update_num)
							inserted_at = fixed_update_idx
						} else {
							if inserted_at > fixed_update_idx {
								fixed_update = append(fixed_update[:inserted_at], fixed_update[inserted_at+1:]...)
								fixed_update = slices.Insert(fixed_update, fixed_update_idx, incorrect_update_num)
								inserted_at = fixed_update_idx
							}
						}
					}
				}
			}
			if inserted_at == -1 {
				fixed_update = append(fixed_update, incorrect_update_num)
			}
		}
		fixed_updates = append(fixed_updates, fixed_update)
	}
	return fixed_updates
}

func SumMiddle(updates [][]int) int {
	sum := 0
	for _, update := range updates {
		sum += update[len(update)/2]
	}
	return sum
}

func VerifyUpdates(updates [][]int, rules map[int][]int) ([][]int, [][]int) {
	correct_updates := [][]int{}
	incorrect_updates := [][]int{}
	for _, update := range updates {
		correct := true
	update_check_loop:
		for idx_update_num := len(update) - 1; idx_update_num > -1; idx_update_num-- {
			for _, rule := range rules[update[idx_update_num]] {
				for idx_num_verif := 0; idx_num_verif < idx_update_num; idx_num_verif++ {
					if rule == update[idx_num_verif] {
						correct = false
						break update_check_loop
					}
				}
			}
		}
		if correct {
			correct_updates = append(correct_updates, update)
		} else {
			incorrect_updates = append(incorrect_updates, update)
		}
	}
	return correct_updates, incorrect_updates
}

func FormatInput(input string) (map[int][]int, [][]int) {
	rules := make(map[int][]int)
	updates := [][]int{}
	reading_rules := true
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			reading_rules = false
			continue
		}
		if reading_rules {
			rule := strings.Split(line, "|")
			key, _ := strconv.Atoi(rule[0])
			value, _ := strconv.Atoi(rule[1])
			rules[key] = append(rules[key], value)
		} else {
			update := []int{}
			update_numbers := strings.Split(line, ",")
			for _, num := range update_numbers {
				converted_num, _ := strconv.Atoi(num)
				update = append(update, converted_num)
			}
			updates = append(updates, update)
		}
	}
	return rules, updates
}

func ReadInput(path string) string {
	content, _ := os.ReadFile(path)
	return string(content)
}
