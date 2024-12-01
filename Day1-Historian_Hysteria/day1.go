package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	list_input := ReadInput("./input-day1-historian_hysteria.txt")

	list_location_1, list_location_2 := GetFormatedListsLocID(list_input)

	fmt.Printf("Total distance: %v\n", GetTotalDistance(list_location_1, list_location_2))

	fmt.Printf("Similarity score: %v\n", GetSimilarityScore(list_location_1, list_location_2))
}

func GetSimilarityScore(list_location_1 []int, list_location_2 []int) int {
	var similarity_score int
	for i1 := range list_location_1 {
		var similarity int
		for i2 := range list_location_2 {
			if list_location_1[i1] == list_location_2[i2] {
				similarity++
			}
		}
		location_similarity_score := list_location_1[i1] * similarity
		similarity_score += location_similarity_score
	}
	return similarity_score
}

func GetTotalDistance(list_location_1 []int, list_location_2 []int) int {
	var total_distance int
	for i := range list_location_1 {
		if list_location_1[i] < list_location_2[i] {
			total_distance += list_location_2[i] - list_location_1[i]
		} else {
			total_distance += list_location_1[i] - list_location_2[i]
		}
	}
	return total_distance
}

func GetFormatedListsLocID(list_input []string) ([]int, []int) {
	list_location_int, err := ConvStringsToIntegers(list_input)
	if err != nil {
		log.Fatalf("Error converting input file from string to int: %v", err)
	}
	var list_location_1 []int
	var list_location_2 []int
	for i, location := range list_location_int {
		if i%2 == 0 {
			list_location_1 = append(list_location_1, location)
		} else {
			list_location_2 = append(list_location_2, location)
		}
	}
	sort.Ints(list_location_1)
	sort.Ints(list_location_2)
	return list_location_1, list_location_2
}

func ReadInput(path string) []string {
	input, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}
	list_input := strings.Fields(string(input))
	return list_input
}

func ConvStringsToIntegers(list_input []string) ([]int, error) {
	list_loc_int := make([]int, len(list_input))
	for i, location := range list_input {
		loc_int, err := strconv.Atoi(location)
		if err != nil {
			return nil, err
		}
		list_loc_int[i] = loc_int
	}
	return list_loc_int, nil
}
